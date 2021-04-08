package rules

import (
	"errors"
	"fmt"
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"reflect"
	"regexp"
)

type Regex struct {}

func (rule Regex) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s format is not valid", field)
}

func (rule Regex) Check(config types.RuleConfig) error {
	if len(config.RuleArgs) == 0 {
		return errors.New("missing args for regex validation")
	}
	// Don't check nil value
	if config.FieldValue == nil {
		return nil
	}
	err := utils.ErrorMsg(config.ErrMsg, rule.GetError(0, config.FieldName, config.RuleArgs))

	regex, regErr := regexp.Compile(config.RuleArgs[0])

	if regErr != nil {
		return fmt.Errorf("regex args is not valid: %s", regErr.Error())
	}

	value := fmt.Sprint(config.FieldValue)

	if value == "" {
		return nil
	}

	if !regex.MatchString(value) {
		return err
	}

	return nil
}
