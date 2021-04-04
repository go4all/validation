package rules

import (
	"testing"
)

var required Required

func TestRequired(t *testing.T) {
	t.Run("Test with string", func(t *testing.T) {
		value := "hello"
		err := required.Check("Greet", value, []string{}, "")
		if err != nil {
			t.Error("Error was not expected")
		}
	})

	t.Run("Test with empty string", func(t *testing.T) {
		value := ""
		err := required.Check("Greet", value, []string{}, "")
		if err == nil {
			t.Error("Error was expected")
		}
	})

	t.Run("Test with slice", func(t *testing.T) {
		value := []string{"Abu Bakkar", "Siddique"}
		err := required.Check("List", value, []string{}, "")
		if err != nil {
			t.Error("Error was not expected")
		}
	})

	t.Run("Test with empty slice", func(t *testing.T) {
		var value []string
		err := required.Check("List", value, []string{}, "")
		if err == nil {
			t.Error("Error was expected")
		}
	})

	t.Run("Test with nil slice", func(t *testing.T) {
		var value []string = nil
		err := required.Check("List", value, []string{}, "")
		if err == nil {
			t.Error("Error was expected")
		}
	})

	t.Run("Test with struct", func(t *testing.T) {
		value := struct{
			Username string
		}{
			Username: "john_do",
		}
		err := required.Check("User", value, []string{}, "")
		if err != nil {
			t.Error("Error was not expected")
		}
	})

	t.Run("Test with empty struct", func(t *testing.T) {
		value := struct{}{}
		err := required.Check("User", value, []string{}, "")
		if err == nil {
			t.Error("Error was expected")
		}
	})
}
