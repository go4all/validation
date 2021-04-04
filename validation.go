package validation

import (
	"fmt"
	"regexp"
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
func Run(request CanValidate) (bool, ErrorBag) {
	errs := make(ErrorBag)
	_rules, _messages := request.Validation()

	for fieldPath, fieldRules := range _rules {
		// Validate fieldPath
		regex := regexp.MustCompile(`^[A-z]\w*(\.[A-z]\w*)*$`)
		if !regex.MatchString(fieldPath) {
			panic(fmt.Sprintf("'%s' field path is not valid", fieldPath))
		}

		fieldErrors := make([]string, 0)
		for _, rule := range fieldRules {
			if rule == "" {
				continue
			}
			ruleName, ruleArgs := ParseRule(rule)
			ruleCheck := GetRuleCheck(ruleName)

			if ruleCheck == nil {
				panic("'" + ruleName + "' rule is missing")
			}

			fieldName, fieldValue := ValueByJsonTag(request, fieldPath)

			err := ruleCheck(fieldName, fieldValue, ruleArgs, _messages[fieldPath][ruleName])
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