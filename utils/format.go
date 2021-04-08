package utils

import (
	"fmt"
	"regexp"
)

// ValidateFieldPath should valid field path
// Valid Format "field[.subField]"
func ValidateFieldPath(fieldPath string) error {
	regex := regexp.MustCompile(`^[A-z]\w*(\.[A-z]\w*)*$`)
	if !regex.MatchString(fieldPath) {
		return fmt.Errorf("'%s' field path is not valid", fieldPath)
	}
	return nil
}

// ValidateRule should validate rule format
// Valid Format: "name[:arg[,arg]]"
func ValidateRule(rule string) error {
	regex := regexp.MustCompile(`^\w+(:(\w+|-?\d+(\.\d+)?)(,(\w+|-?\d+(\.\d+)?))*)?$`)
	if !regex.MatchString(rule) {
		return fmt.Errorf("'%s' rule is not valid", rule)
	}
	return nil
}
