package types

import "github.com/spectrocloud-labs/validator/api/v1alpha1"

// ValidationRuleResult is the result of the execution of a validation rule by a validator
type ValidationRuleResult struct {
	Condition *v1alpha1.ValidationCondition
	State     *v1alpha1.ValidationState
}

type SinkType string

const (
	SinkTypeAlertmanager SinkType = "alertmanager"
	SinkTypeSlack        SinkType = "slack"
)
