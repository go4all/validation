package rules

import (
	"fmt"
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"reflect"
	"regexp"
)

type AlphaNum struct{}

func (rule AlphaNum) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s format is not valid", field)
}

func (rule AlphaNum) Check(config types.RuleConfig) error {
	// Don't check nil value
	if config.FieldValue == nil {
		return nil
	}
	err := utils.ErrorMsg(config.ErrMsg, rule.GetError(0, config.FieldName, config.RuleArgs))

	value, ok := config.FieldValue.(string)

	if !ok {
		return err
	}

	if value == "" {
		return nil
	}

	regex := regexp.MustCompile(`^[A-z0-9_]+$`)

	if !regex.MatchString(value) {
		return err
	}

	return nil
}
