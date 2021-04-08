package validation

import (
	"github.com/go4all/validation/rules"
	"github.com/go4all/validation/types"
)

// ruleList is a map of all RuleCheck functions that can be used in request validation.
var ruleList = make(types.RuleList)

func init() {
	AddRuleCheck("max", rules.Max{}.Check)
	AddRuleCheck("min", rules.Min{}.Check)
	AddRuleCheck("email", rules.Email{}.Check)
	AddRuleCheck("match", rules.Match{}.Check)
	AddRuleCheck("regex", rules.Regex{}.Check)
	AddRuleCheck("required", rules.Required{}.Check)
	AddRuleCheck("alpha_num", rules.AlphaNum{}.Check)
}

// AddRuleCheck will add RuleCheck in the map of RuleCheck functions with provided name to retrieve later
// Use this to add custom RuleCheck function
func AddRuleCheck(name string, rule types.RuleCheck) {
	ruleList[name] = rule
}
// GetRuleCheck will return a RuleCheck function added with provided name from the map of available RuleCheck functions
func GetRuleCheck(name string) types.RuleCheck {
	rule, exists := ruleList[name]
	if !exists {
		return nil
	}
	return rule
}
