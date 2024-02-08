package validationresult

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/pkg/constants"
	"github.com/spectrocloud-labs/validator/pkg/types"
	"github.com/spectrocloud-labs/validator/pkg/util"
)

const validationErrorMsg = "Validation failed with an unexpected error"

// HandleExistingValidationResult processes a preexisting validation result for the active validator
func HandleExistingValidationResult(nn ktypes.NamespacedName, vr *v1alpha1.ValidationResult, l logr.Logger) {
	switch vr.Status.State {

	case v1alpha1.ValidationInProgress:
		// validations are only left in progress if an unexpected error occurred
		l.V(0).Info("Previous validation failed with unexpected error", "name", nn.Name, "namespace", nn.Namespace)

	case v1alpha1.ValidationFailed:
		// log validation failure, but continue and retry
		cs := getInvalidConditions(vr.Status.Conditions)
		if len(cs) > 0 {
			for _, c := range cs {
				l.V(0).Info(
					"Validation failed. Retrying.", "name", nn.Name, "namespace", nn.Namespace,
					"validation", c.ValidationRule, "error", c.Message, "details", c.Details, "failures", c.Failures,
				)
			}
		}

	case v1alpha1.ValidationSucceeded:
		// log validation success, continue to re-validate
		l.V(0).Info("Previous validation succeeded. Re-validating.", "name", nn.Name, "namespace", nn.Namespace)
	}
}

// HandleNewValidationResult creates a new validation result for the active validator
func HandleNewValidationResult(c client.Client, vr *v1alpha1.ValidationResult, l logr.Logger) error {

	// Create the ValidationResult
	if err := c.Create(context.Background(), vr, &client.CreateOptions{}); err != nil {
		l.V(0).Error(err, "failed to create ValidationResult", "name", vr.Name, "namespace", vr.Namespace)
		return err
	}

	var err error
	nn := ktypes.NamespacedName{Name: vr.Name, Namespace: vr.Namespace}

	// Update the ValidationResult's status
	for i := 0; i < constants.StatusUpdateRetries; i++ {
		if err := c.Get(context.Background(), nn, vr); err != nil {
			l.V(0).Error(err, "failed to get ValidationResult", "name", nn.Name, "namespace", nn.Namespace)
			return err
		}
		vr.Status = v1alpha1.ValidationResultStatus{State: v1alpha1.ValidationInProgress}
		err = c.Status().Update(context.Background(), vr)
		if err == nil {
			return nil
		}
		l.V(1).Info("warning: failed to update ValidationResult status", "name", vr.Name, "namespace", vr.Namespace, "error", err.Error())
	}
	return err
}

// SafeUpdateValidationResult updates the overall validation result, ensuring
// that the overall validation status remains failed if a single rule fails
func SafeUpdateValidationResult(c client.Client, nn ktypes.NamespacedName, res *types.ValidationResult, resCount int, resErr error, l logr.Logger) {
	var err error
	var updated bool
	ctx := context.Background()
	vr := &v1alpha1.ValidationResult{}

	for i := 0; i < constants.StatusUpdateRetries; i++ {
		if err := c.Get(ctx, nn, vr); err != nil {
			l.V(0).Error(err, "failed to get ValidationResult", "name", nn.Name, "namespace", nn.Namespace)
			continue
		}
		vr.Spec.ExpectedResults = resCount
		if err := c.Update(ctx, vr); err != nil {
			l.V(1).Info("warning: failed to update ValidationResult", "name", nn.Name, "namespace", nn.Namespace, "error", err.Error())
			continue
		}
		l.V(0).Info("Updated ValidationResult", "plugin", vr.Spec.Plugin, "expectedResults", vr.Spec.ExpectedResults)
		updated = true
		break
	}
	if !updated {
		l.V(0).Error(err, "failed to update ValidationResult", "name", nn.Name, "namespace", nn.Namespace)
		return
	}

	for i := 0; i < constants.StatusUpdateRetries; i++ {
		if err := c.Get(ctx, nn, vr); err != nil {
			l.V(0).Error(err, "failed to get ValidationResult", "name", nn.Name, "namespace", nn.Namespace)
			continue
		}
		updateValidationResultStatus(vr, res, resErr)
		if err := c.Status().Update(ctx, vr); err != nil {
			l.V(1).Info("warning: failed to update ValidationResult status", "name", nn.Name, "namespace", nn.Namespace, "error", err.Error())
			continue
		}
		l.V(0).Info(
			"Updated ValidationResult status", "state", res.State, "reason", res.Condition.ValidationRule,
			"message", res.Condition.Message, "details", res.Condition.Details,
			"failures", res.Condition.Failures, "time", res.Condition.LastValidationTime,
		)
		return
	}

	l.V(0).Error(err, "failed to update ValidationResult status", "name", nn.Name, "namespace", nn.Namespace)
}

// updateValidationResultStatus updates a ValidationResult's status with the result of a single validation rule
func updateValidationResultStatus(vr *v1alpha1.ValidationResult, res *types.ValidationResult, resErr error) {

	// Finalize result State and Condition in the event of an unexpected error
	if resErr != nil {
		res.State = util.Ptr(v1alpha1.ValidationFailed)
		res.Condition.Status = corev1.ConditionFalse
		res.Condition.Message = validationErrorMsg
		res.Condition.Failures = append(res.Condition.Failures, resErr.Error())
	}

	// Update and/or insert the ValidationResult's Conditions with the latest Condition
	idx := getConditionIndexByValidationRule(vr.Status.Conditions, res.Condition.ValidationRule)
	if idx == -1 {
		vr.Status.Conditions = append(vr.Status.Conditions, *res.Condition)
	} else {
		vr.Status.Conditions[idx] = *res.Condition
	}

	// Set State to:
	// - ValidationFailed if ANY condition failed
	// - ValidationSucceeded if ALL conditions succeeded
	// - ValidationInProgress otherwise
	vr.Status.State = *res.State
	for _, c := range vr.Status.Conditions {
		if c.Status == corev1.ConditionTrue {
			vr.Status.State = v1alpha1.ValidationSucceeded
		}
		if c.Status == corev1.ConditionFalse {
			vr.Status.State = v1alpha1.ValidationFailed
			break
		}
	}
}

// getInvalidConditions filters a ValidationCondition array and returns all conditions corresponding to a failed validation
func getInvalidConditions(conditions []v1alpha1.ValidationCondition) []v1alpha1.ValidationCondition {
	invalidConditions := make([]v1alpha1.ValidationCondition, 0)
	for _, c := range conditions {
		if strings.HasPrefix(c.ValidationRule, constants.ValidationRulePrefix) && c.Status == corev1.ConditionFalse {
			invalidConditions = append(invalidConditions, c)
		}
	}
	return invalidConditions
}

// getConditionIndexByValidationRule retrieves the index of a condition from a ValidationCondition array matching a specific reason
func getConditionIndexByValidationRule(conditions []v1alpha1.ValidationCondition, validationRule string) int {
	for i, c := range conditions {
		if c.ValidationRule == validationRule {
			return i
		}
	}
	return -1
}
