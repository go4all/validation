package rules

import (
	"testing"
)

var min Min

func TestMin(t *testing.T) {
	t.Run("Test with string within min length", func(t *testing.T) {
		value := "hello"
		err := min.Check("Greet", value, []string{"5"}, "")
		if err != nil {
			t.Error("Error was not expected")
		}
	})
}
