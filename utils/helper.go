package utils

import (
	"strings"
)

func GetFieldNameFromPath(path string) (string, bool) {
	err := ValidateFieldPath(path)
	if err != nil {
		return "", false
	}

	segments := strings.Split(path, ".")

	return segments[len(segments)-1], true
}
