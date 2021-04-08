package rules

import (
	"github.com/go4all/validation/types"
	"testing"
)

var match Match

func TestMatch(t *testing.T) {
	t.Run("Test with custom message", func(t *testing.T) {
		customMsg := "Your values are different"
		err := match.Check(types.RuleConfig{
			FieldName:  "password",
			FieldValue: "secret1",
			ErrMsg:     customMsg,
			RuleArgs:   []string{"password"},
			Values: map[string]interface{}{
				"password": "secret",
			},
		})
		if err == nil {
			t.Error("Expected error got nil")
		} else if err.Error() != customMsg {
			t.Error("Expected custom error got default error")
		}
	})

	t.Run("Test with matching values", func(t *testing.T) {
		err := match.Check(types.RuleConfig{
			FieldName:  "password",
			FieldValue: "secret",
			RuleArgs:   []string{"password"},
			Values: map[string]interface{}{
				"password": "secret",
			},
		})
		if err != nil {
			t.Error("Expected nil got error")
		}
	})

	t.Run("Test with different values", func(t *testing.T) {
		err := match.Check(types.RuleConfig{
			FieldName:  "password",
			FieldValue: "secret1",
			RuleArgs:   []string{"password"},
			Values: map[string]interface{}{
				"password": "secret",
			},
		})
		if err == nil {
			t.Error("Expected error got nil")
		}
	})

	t.Run("Test with different values types", func(t *testing.T) {
		err := match.Check(types.RuleConfig{
			FieldName:  "age",
			FieldValue: 15,
			RuleArgs:   []string{"age"},
			Values: map[string]interface{}{
				"age": "15",
			},
		})
		if err == nil {
			t.Error("Expected error got nil")
		}
	})

	t.Run("Test with missing field value", func(t *testing.T) {
		err := match.Check(types.RuleConfig{
			FieldName:  "password",
			FieldValue: nil,
			RuleArgs:   []string{"password"},
		})
		if err != nil {
			t.Error("Expected nil got error")
		}
	})

	t.Run("Test with missing matching value", func(t *testing.T) {
		err := match.Check(types.RuleConfig{
			FieldName:  "password",
			FieldValue: "secret",
			RuleArgs:   []string{"password"},
		})
		if err == nil {
			t.Error("Expected error got nil")
		}
	})
}
