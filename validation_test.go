package validation

import (
	"testing"

	"github.com/go4all/validation/types"
)

type Job struct {
	Title   string `json:"title"`
	Company string `json:"company"`
}

type TestRequest struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
	Job    Job      `json:"job"`
}

func (tr TestRequest) Validation() (types.RuleMap, types.MessageMap) {
	rules := types.RuleMap{
		"name":        {"required", "max:24"},
		"age":         {"required", "min:18"},
		"skills":      {"required", "max:53"},
		"job.title":   {"required"},
		"job.company": {"required"},
	}

	messages := types.MessageMap{
		"name": map[string]string{
			"required": "Please fill in your name",
			"max":      "Your name is too long",
		},
		"skills": {
			"required": "Skills are required",
		},
	}
	return rules, messages
}

func TestValidation_Run(t *testing.T) {
	t.Run("Test validation", func(t *testing.T) {
		request := TestRequest{
			Name:   "Abu Bakkar Siddique",
			Age:    21,
			Skills: []string{"JavaScript", "HTML", "CSS", "React", "Vue"},
			Job: Job{
				Title: "Software Engineer",
			},
		}
		errs, err := Run(request)

		if err != nil {
			t.Error("Got error while validation data")
		}

		ruleErr, ok := errs["job.company"]

		if !ok {
			t.Error("'job.company' key missing from errors")
		}

		if len(ruleErr) == 0 {
			t.Error("Expected error for 'job.company'")
		}

		if ruleErr[0] != "company is required" {
			t.Errorf("Expected error 'company is required' got '%s'", ruleErr[0])
		}
	})
}
