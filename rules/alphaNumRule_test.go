package rules

import (
	"github.com/go4all/validation/types"
	"testing"
)

var alphaNum AlphaNum

func TestAlphaNum(t *testing.T) {
	t.Run("Test with custom message", func(t *testing.T) {
		customMsg := "Username is not valid"
		err := alphaNum.Check(types.RuleConfig{
			FieldName: "username",
			FieldValue: "secret 23",
			ErrMsg: customMsg,
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
			"_secret123_",
			"Secret123_F",
		}

		for _, value := range values {
			err := alphaNum.Check(types.RuleConfig{
				FieldName: "username",
				FieldValue: value,
			})

			if err != nil {
				t.Errorf("Expected nil for value '%s' got error", value)
			}
		}
	})

	t.Run("Test with invalid values", func(t *testing.T) {
		values := []string{
			"!secret",
			"secret-123",
			"secret 123",
		}

		for _, value := range values {
			err := alphaNum.Check(types.RuleConfig{
				FieldName: "username",
				FieldValue: value,
			})

			if err == nil {
				t.Errorf("Expected error for value '%s' got nil", value)
			}
		}
	})
}
