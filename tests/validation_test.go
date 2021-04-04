package tests

import (
	"fmt"
	"github.com/go4all/validaiton"
	"testing"
)

type Job struct {
	Title string `json:"title"`
	Company string `json:"company"`
}

type TestRequest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skills []string `json:"skills"`
	Job Job `json:"job"`
}

func (tr TestRequest) Validation() (validation.RuleMap, validation.MessageMap)  {
	rules := validation.RuleMap{
		"name": {"required", "max:24"},
		"age": {"required", "min:18"},
		"skills": {"required", "max:5"},
		"job.title": {"required"},
		"job.company": {"required"},
	}

	messages := validation.MessageMap{
		"name": map[string]string{
			"required": "Please fill in your name",
			"max": "Your name is too long",
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
			Name: "Abu Bakkar Siddique",
			Age: 21,
			Skills: []string{"JavaScript", "HTML", "CSS", "React", "Vue"},
			Job: Job{
				Title: "Software Engineer",
			},
		}
		valid, errs := validation.Run(request)
		fmt.Println(valid, errs)
	})
}
