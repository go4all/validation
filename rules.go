package validation

// ruleList is a map of all RuleCheck functions that can be used in request validation.
var ruleList = make(RuleList)

// AddRuleCheck will add RuleCheck in the map of RuleCheck functions with provided name to retrieve later
// Use this to add custom RuleCheck function
func AddRuleCheck(name string, rule RuleCheck) {
	ruleList[name] = rule
}
// GetRuleCheck will return a RuleCheck function added with provided name from the map of available RuleCheck functions
func GetRuleCheck(name string) RuleCheck {
	rule, exists := ruleList[name]
	if !exists {
		return nil
	}
	return rule
}
