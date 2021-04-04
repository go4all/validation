package tests

import (
	"github.com/apni-market/api/app/validation"
	"testing"
)

var max validation.Max

func TestMax(t *testing.T) {
	t.Run("Test with string within max length", func(t *testing.T) {
		value := "hello"
		err := max.Check("Greet", value, []string{"5"}, "")
		if err != nil {
			t.Error("Error was not expected")
		}
	})
}
