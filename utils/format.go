package utils

import (
	"fmt"
	"regexp"
)

func ValidateFieldPath(fieldPath string) error {
	regex := regexp.MustCompile(`^[A-z]\w*(\.[A-z]\w*)*$`)
	if !regex.MatchString(fieldPath) {
		return fmt.Errorf("'%s' field path is not valid", fieldPath)
	}
	return nil
}

func ValidateRule(rule string) error {
	regex := regexp.MustCompile(`^\w+(:\w+(,\w+)*)?$`)
	if !regex.MatchString(rule) {
		return fmt.Errorf("'%s' rule is not valid", rule)
	}
	return nil
}
