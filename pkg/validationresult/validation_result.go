package validationresult

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
	"github.com/spectrocloud-labs/validator/pkg/constants"
	"github.com/spectrocloud-labs/validator/pkg/types"
	"github.com/spectrocloud-labs/validator/pkg/util/ptr"
)

// HandleExistingValidationResult processes a preexisting validation result for the active validator
func HandleExistingValidationResult(nn ktypes.NamespacedName, vr *v1alpha1.ValidationResult, l logr.Logger) (*ctrl.Result, error) {
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

	return nil, nil
}

// HandleNewValidationResult creates a new validation result for the active validator
func HandleNewValidationResult(c client.Client, plugin string, nn ktypes.NamespacedName, vr *v1alpha1.ValidationResult, l logr.Logger) (*ctrl.Result, error) {

	// Create the ValidationResult
	vr.ObjectMeta = metav1.ObjectMeta{
		Name:      nn.Name,
		Namespace: nn.Namespace,
	}
	vr.Spec = v1alpha1.ValidationResultSpec{
		Plugin: plugin,
	}
	if err := c.Create(context.Background(), vr, &client.CreateOptions{}); err != nil {
		l.V(0).Error(err, "failed to create ValidationResult", "name", nn.Name, "namespace", nn.Namespace)
		return &ctrl.Result{}, err
	}

	// Update the ValidationResult's status
	vr.Status = v1alpha1.ValidationResultStatus{
		State: v1alpha1.ValidationInProgress,
	}
	if err := c.Status().Update(context.Background(), vr); err != nil {
		l.V(0).Error(err, "failed to update ValidationResult status", "name", nn.Name, "namespace", nn.Namespace)
		return &ctrl.Result{}, err
	}

	return nil, nil
}

// SafeUpdateValidationResult updates the overall validation result, ensuring
// that the overall validation status remains failed if a single rule fails
func SafeUpdateValidationResult(
	c client.Client, nn ktypes.NamespacedName, res *types.ValidationResult,
	failed *types.MonotonicBool, err error, l logr.Logger,
) {
	if err != nil {
		res.State = ptr.Ptr(v1alpha1.ValidationFailed)
		res.Condition.Status = corev1.ConditionFalse
		res.Condition.Message = "Validation failed with an unexpected error"
		res.Condition.Failures = append(res.Condition.Failures, err.Error())
	}

	didFail := *res.State == v1alpha1.ValidationFailed
	failed.Update(didFail)
	if failed.Ok && !didFail {
		res.State = ptr.Ptr(v1alpha1.ValidationFailed)
	}

	if err := updateValidationResult(c, nn, res, l); err != nil {
		l.V(0).Error(err, "failed to update ValidationResult")
	}
}

// updateValidationResult updates the ValidationResult for the active validation rule
func updateValidationResult(c client.Client, nn ktypes.NamespacedName, res *types.ValidationResult, l logr.Logger) error {
	vr := &v1alpha1.ValidationResult{}
	if err := c.Get(context.Background(), nn, vr); err != nil {
		return fmt.Errorf("failed to get ValidationResult %s in namespace %s: %v", nn.Name, nn.Namespace, err)
	}
	vr.Status.State = *res.State

	idx := getConditionIndexByValidationRule(vr.Status.Conditions, res.Condition.ValidationRule)
	if idx == -1 {
		vr.Status.Conditions = append(vr.Status.Conditions, *res.Condition)
	} else {
		vr.Status.Conditions[idx] = *res.Condition
	}

	if err := c.Status().Update(context.Background(), vr); err != nil {
		l.V(0).Error(err, "failed to update ValidationResult")
		return err
	}
	l.V(0).Info(
		"Updated ValidationResult", "state", res.State, "reason", res.Condition.ValidationRule,
		"message", res.Condition.Message, "details", res.Condition.Details,
		"failures", res.Condition.Failures, "time", res.Condition.LastValidationTime,
	)

	return nil
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
