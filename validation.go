package validation

import (
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"strings"
)

// ParseRule will parse rule into processable format
// Input: "range:20,30"
// Output: ("range", []string{"20", "30"})
func ParseRule(rule string) (string, []string) {
	segments := strings.Split(rule, ":")
	if len(segments) == 1 {
		return segments[0], []string{}
	}
	return segments[0], strings.Split(segments[1], ",")
}

// Run will validate request with provided validation rules and return error messages if validation fails
func Run(request types.CanValidate) (bool, types.ErrorBag) {
	errs := make(types.ErrorBag)
	_rules, _messages := request.Validation()
	_values := GetValues(request)

	for fieldPath, fieldRules := range _rules {
		// Validate fieldPath
		err := utils.ValidateFieldPath(fieldPath)
		if err != nil {
			panic(err)
		}
		// This will hold errors for a specific field
		fieldErrors := make([]string, 0)

		for _, rule := range fieldRules {
			// Empty string for a rule should be ignored
			err = utils.ValidateRule(rule)
			if err != nil {
				panic(err)
			}
			ruleName, ruleArgs := ParseRule(rule)
			ruleCheck := GetRuleCheck(ruleName)

			if ruleCheck == nil {
				panic("'" + ruleName + "' rule is missing")
			}

			fieldName, fieldValue := ValueByFieldPath(_values, fieldPath)

			ruleConfig := types.RuleConfig{
				FieldName: fieldName,
				FieldValue: fieldValue,
				RuleArgs: ruleArgs,
				ErrMsg: _messages[fieldPath][ruleName],
				Values: _values,
			}

			err := ruleCheck(ruleConfig)
			if err != nil {
				fieldErrors = append(fieldErrors, err.Error())
			}
		}
		if len(fieldErrors) > 0 {
			errs[fieldPath] = fieldErrors
		}
	}

	return len(errs) == 0, errs
}