package rules

import (
	"errors"
	"fmt"
	"github.com/go4all/validation/types"
	"github.com/go4all/validation/utils"
	"reflect"
)

type Match struct {}

func (rule Match) GetError(kind reflect.Kind, field string, args []string) string {
	return fmt.Sprintf("%s should match with %s", field, args[0])
}

func (rule Match) Check(config types.RuleConfig) error {
	if len(config.RuleArgs) == 0 {
		return errors.New("missing args for match validation")
	}
	// Don't check nil value
	if config.FieldValue == nil {
		return nil
	}
	err := utils.ErrorMsg(config.ErrMsg, rule.GetError(0, config.FieldName, config.RuleArgs))

	matchingField := config.RuleArgs[0]

	matchingValue, ok := config.Values[matchingField]

	if !ok || matchingValue != config.FieldValue{
		return err
	}

	return nil
}
