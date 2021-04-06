package utils

import (
	"testing"
)

func TestGetFieldNameFromPath(t *testing.T) {
	t.Run("Test with valid paths", func(t *testing.T) {
		paths := []string{
			"profile",
			"user.email",
			"user.profile.firstname",
		}

		for _, path := range paths {
			field, ok := GetFieldNameFromPath(path)

			if !ok || field == "" {
				t.Error("Expected field name got empty string")
			}
		}
	})

	t.Run("Test with invalid paths", func(t *testing.T) {
		paths := []string{
			".",
			"user.",
			".email",
		}

		for _, path := range paths {
			field, ok := GetFieldNameFromPath(path)

			if ok || field != "" {
				t.Error("Expected error got field name")
			}
		}
	})
}
