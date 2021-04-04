package types

import "reflect"

// RuleConfig
type RuleConfig struct {
	FieldName string
	FieldValue interface{}
	RuleArgs []string
	ErrMsg string
	Values map[string]interface{}
}
// ErrorBag is map of error messages returned from request validation
type ErrorBag map[string][]string
// CanValidate is an interface that defines a struct to be validatable
type CanValidate interface {
	Validation() (RuleMap, MessageMap)
}
// Rule is an interface that defines a struct to be used as rule
type Rule interface {
	GetError(kind reflect.Kind, field string, args []string) string
	Check(RuleConfig) error
}
// RuleCheck is a function responsible for validating value depending on rule type
type RuleCheck func(config RuleConfig) error
// RuleList is a map of all registered rules
type RuleList map[string]RuleCheck
// RuleMap is a map of rules required to validate a request
type RuleMap map[string][]string
// MessageMap is a map of custom error messages for validation rules
type MessageMap map[string]map[string]string
