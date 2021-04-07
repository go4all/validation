package rules

import (
	"fmt"
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"reflect"
	"regexp"
)

type Email struct {}

func (rule Email) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s is not a valid email address", field)
}

func (rule Email) Check(config types.RuleConfig) error {
	// Don't check nil value
	if config.FieldValue == nil {
		return nil
	}
	err := utils.ErrorMsg(config.ErrMsg, rule.GetError(0, config.FieldName, config.RuleArgs))

	email, ok := config.FieldValue.(string)

	if !ok {
		return err
	}

	if email == "" {
		return nil
	}

	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !regex.MatchString(email) {
		return err
	}

	return nil
}
