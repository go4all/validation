package rules

import (
	"github.com/go4all/validation/types"
	"testing"
)

var email Email

func TestEmail(t *testing.T) {
	t.Run("Test with custom message", func(t *testing.T) {
		customMsg := "Email is not valid"
		err := email.Check(types.RuleConfig{
			FieldName:  "email",
			FieldValue: "test_example.com",
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
			"hello@example",
			"hello+12@example.com",
			"test@example.com",
			"person.one@test.com",
			"hello123@test.com",
		}

		for _, value := range values {
			err := email.Check(types.RuleConfig{
				FieldName:  "email",
				FieldValue: value,
			})

			if err != nil {
				t.Errorf("Expected nil for value '%s' got error", value)
			}
		}
	})

	t.Run("Test with invalid values", func(t *testing.T) {
		values := []string{
			"hello+1test.com",
			"test-123@example.",
			"@test.com",
		}

		for _, value := range values {
			err := email.Check(types.RuleConfig{
				FieldName:  "email",
				FieldValue: value,
			})

			if err == nil {
				t.Errorf("Expected error for value '%s' got nil", value)
			}
		}
	})
}
