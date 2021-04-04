package validation

import (
	"errors"
	"fmt"
)
// ErrorMsg will return formatted error message if custom message is empty string
func ErrorMsg(custom, format string, values ...interface{}) error {
	defaultMsg := fmt.Sprintf(format, values...)

	actualMsg := custom

	if actualMsg == "" {
		actualMsg = defaultMsg
	}

	return errors.New(actualMsg)
}
