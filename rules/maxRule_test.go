package rules

import (
	"github.com/go4all/validation/types"
	"testing"
)

var max Max

func TestMax(t *testing.T) {
	t.Run("Test with string within max length", func(t *testing.T) {
		value := "hello"
		err := max.Check(types.RuleConfig{
			FieldName: "Greet",
			FieldValue: value,
			RuleArgs: []string{"5"},
		})
		if err != nil {
			t.Error("Error was not expected")
		}
	})
}
