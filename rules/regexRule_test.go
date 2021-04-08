package rules

import (
	"github.com/go4all/validation/types"
	"testing"
)

var regex Regex

func TestRegex(t *testing.T) {
	t.Run("Test with custom message", func(t *testing.T) {
		customMsg := "Username is not valid"
		err := regex.Check(types.RuleConfig{
			FieldName:  "username",
			FieldValue: "ads-asf",
			RuleArgs:   []string{"^[a-z][a-z0-9_]{5,23}$"},
			ErrMsg:     customMsg,
		})
		if err == nil {
			t.Error("Expected error got nil")
		} else if err.Error() != customMsg {
			t.Error("Expected custom error got default error")
		}
	})

	t.Run("Test with valid values", func(t *testing.T) {
		values := []string{
			"secret",
			"secret123",
			"secret123_",
			"secret123_f",
		}

		for _, value := range values {
			err := regex.Check(types.RuleConfig{
				FieldName:  "username",
				FieldValue: value,
				RuleArgs:   []string{"^[a-z][a-z0-9_]{5,23}$"},
			})

			if err != nil {
				t.Errorf("Expected nil for value '%s' got error", value)
			}
		}
	})

	t.Run("Test with invalid values", func(t *testing.T) {
		values := []string{
			"!secret",
			"Science",
			"secret-123",
			"secret 123",
		}

		for _, value := range values {
			err := regex.Check(types.RuleConfig{
				FieldName:  "username",
				FieldValue: value,
				RuleArgs:   []string{"^[a-z][a-z0-9_]{5,23}$"},
			})

			if err == nil {
				t.Errorf("Expected error for value '%s' got nil", value)
			}
		}
	})
}
