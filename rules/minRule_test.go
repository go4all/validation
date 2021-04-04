package rules

import (
	"github.com/go4all/validaiton/types"
	"testing"
)

var min Min

func TestMin(t *testing.T) {
	t.Run("Test with string within min length", func(t *testing.T) {
		value := "hello"
		err := min.Check(types.RuleConfig{
			FieldName: "Greet",
			FieldValue: value,
			RuleArgs: []string{"5"},
		})
		if err != nil {
			t.Error("Error was not expected")
		}
	})
}
