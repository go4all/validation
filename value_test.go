package validation

import (
	"testing"
)

type Location struct {
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}

type Profile struct {
	Title string `json:"title"`
	Company string `json:"company"`
	Skills []string `json:"skills"`
	Location Location `json:"location"`
}

type User struct {
	Name string `json:"name"`
	Email string
	Profile Profile `json:"profile"`
}

func TestValueByFieldPath(t *testing.T) {
	t.Run("Test with root field", func(t *testing.T) {
		values := make(map[string]interface{})
		values["name"] = "Abu Bakkar"

		name, value := ValueByFieldPath(values, "name")

		if name != "name" {
			t.Errorf("Expected 'name' got '%s'", name)
		}

		if value == nil {
			t.Error("Expected value got nil")
		}
	})

	t.Run("Test with nest field path", func(t *testing.T) {
		values := make(map[string]interface{})
		values["name"] = "Abu Bakkar"
		values["profile.company"] = "Love,Bonito"

		name, value := ValueByFieldPath(values, "profile.company")

		if name != "company" {
			t.Errorf("Expected 'company' got '%s'", name)
		}

		if value == nil {
			t.Error("Expected value got nil")
		}
	})

	t.Run("Test with missing nest field path", func(t *testing.T) {
		values := make(map[string]interface{})
		values["profile.company"] = "Love,Bonito"

		name, value := ValueByFieldPath(values, "profile.age")

		if name != "profile.age" {
			t.Errorf("Expected 'profile.age' got '%s'", name)
		}

		if value != nil {
			t.Errorf("Expected nil got %s", value)
		}
	})
}

func TestGetValues(t *testing.T) {
	user := User{
		Name: "Abu Bakkar",
		Email: "user@example.com",
		Profile: Profile{
			Title: "Software Engineer",
			Company: "Love,Bonito",
			Skills: []string{"JavaScript", "HTML", "CSS", "Vue", "React"},
			Location: Location{
				City: "Gujrat",
				State: "Punjab",
				Country: "Pakistan",
			},
		},
	}

	output := GetValues(user)

	if _, ok := output["name"]; !ok {
		t.Error("name is missing from output")
	}

	if _, ok := output["email"]; ok {
		t.Error("email should not include in output")
	}

	if _, ok := output["profile.title"]; !ok {
		t.Error("profile.title is missing from output")
	}

	if _, ok := output["profile.location.city"]; !ok {
		t.Error("profile.location.city is missing from output")
	}
}
