This is a validation package for golang web applications inspired by Laravel validation.

###Installation

`go get -u github.com/go4all/validation`

###Usage

This package can be used to validate struct fields. The struct should implement `CanValidate` interface provide by 
this package. Also this package use `json` tag on struct fields.

```go
package example

import "github.com/go4all/validaiton"

type SignUpRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request SignUpRequest) Validation() (validation.RuleMap, validation.MessageMap)  {
    rules := validation.RuleMap{
    	"name": {"required"},
    	"email": {"required", "email"},
    	"password": {"required", "min:8"},
    }
    
    messages := validation.MessageMap{
    	"username": {
    		"required": "Please provide username",
        },
        "password": {
    		"required": "Must provide password",
    		"min": "Password should be at-least 8 characters",
        },
    }
    return rules, messages
}
```