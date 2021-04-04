package utils

import "testing"

func TestValidateRule(t *testing.T) {
	t.Run("Check valid rule", func(t *testing.T) {
		validRules := []string{
			"required",
			"min:12",
			"max:24.3",
			"range:12,24",
		}

		for _, rule := range validRules {
			err := ValidateRule(rule)

			if err != nil {
				t.Errorf("'%s' should be valid", rule)
			}
		}
	})
	t.Run("Check invalid rule", func(t *testing.T) {
		invalidRules := []string{
			"!required",
			"min:",
			"min: 24, 12",
		}

		for _, rule := range invalidRules {
			err := ValidateRule(rule)

			if err == nil {
				t.Errorf("'%s' should be invalid", rule)
			}
		}
	})
}

func TestValidateFieldPath(t *testing.T) {
	t.Run("Check valid field path", func(t *testing.T) {
		validPaths := []string{
			"user",
			"user1",
			"user.name",
			"_user_1._name_",
		}

		for _, path := range validPaths {
			err := ValidateFieldPath(path)

			if err != nil {
				t.Errorf("'%s' should be valid", path)
			}
		}
	})
	t.Run("Check invalid field path", func(t *testing.T) {
		invalidPaths := []string{
			"!user",
			"1user",
			".user",
			"user..name",
			"user.1",
		}

		for _, path := range invalidPaths {
			err := ValidateFieldPath(path)

			if err == nil {
				t.Errorf("'%s' should be invalid", path)
			}
		}
	})
}
