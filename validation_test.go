package validation

import (
	"fmt"
	"github.com/go4all/validation/utils"
	"testing"

	"github.com/go4all/validation/types"
)

type Job struct {
	Title   string `json:"title"`
	Company string `json:"company"`
}

type SignUpRequest struct {
	Name            string   `json:"name"`
	Age             int      `json:"age"`
	Job             Job      `json:"job"`
	Skills          []string `json:"skills"`
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	NewPassword     string   `json:"new_password"`
	ConfirmPassword string   `json:"confirm_password"`
}

func (tr SignUpRequest) Validation() (types.RuleMap, types.MessageMap) {
	rules := types.RuleMap{
		"name":             {"required", "max:24"},
		"age":              {"min:18", "max:65"},
		"skills":           {"required", "max:5"},
		"job.title":        {"required"},
		"email":            {"required", "email"},
		"username":         {"required", "alpha_num"},
		"new_password":     {"required", "min:8"},
		"confirm_password": {"required", "match:new_password"},
	}

	messages := types.MessageMap{
		"name": map[string]string{
			"required": "Let us know your name",
			"max":      "Your name is too long",
		},
		"skills": {
			"required": "Skills are required",
		},
		"confirm_password": {
			"match": "Both passwords should match",
		},
	}
	return rules, messages
}

func TestValidation_Run(t *testing.T) {
	t.Run("Test validation", func(t *testing.T) {
		request := SignUpRequest{
			Name:   "Abu Bakkar Siddique",
			Age:    66,
			Skills: []string{},
			Job: Job{
				Title:   "Software Engineer",
				Company: "Love,Bonito",
			},
			Email:           "hello.com",
			Username:        "234asdf",
			NewPassword:     "secret133",
			ConfirmPassword: "secret123",
		}
		errs, err := Run(request)

		if err != nil {
			t.Error(err.Error())
		}

		fields := map[string][]string{
			"email":            {},
			"confirm_password": {},
		}

		for field, msgs := range fields {
			ruleErrs, ok := errs[field]

			if !ok {
				t.Errorf("Expected errors for '%s'", field)
			}

			if len(ruleErrs) == 0 {
				t.Errorf("Expected errors for '%s', found none", field)
			}

			for _, msg := range msgs {
				if !utils.HasError(msg, ruleErrs) {
					t.Errorf("Missing error message '%s' for '%s'", msg, field)
				}
			}
		}
		fmt.Println("-------------")
		for field, msgs := range errs {
			fmt.Println(field)
			for _, msg := range msgs {
				fmt.Printf("\t- %s\n", msg)
			}
		}
		fmt.Println("-------------")
	})
}
