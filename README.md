This is a validation package for golang web applications inspired by Laravel validation.

### Installation

`go get -u github.com/go4all/validation`

### Usage

You can validate any `struct` type by implement `Validation() (types.RuleMap, types.MessageMap)` method provided by
`CanValidate` interface. 

Following is an example of a `struct` which implemented `CanValidate` interface
```go
package requests

import (
	"github.com/go4all/validation"
	"github.com/go4all/validation/types"
)

type SignUpRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request SignUpRequest) Validation() (types.RuleMap, types.MessageMap)  {
    rules := types.RuleMap{
    	"name": {"required"},
    	"email": {"required", "email"},
    	"new_password": {"required", "min:8"},
    	"confirm_password": {"required", "match:new_password"},
    }
    
    messages := types.MessageMap{
    	"email": {
    		"required": "Please provide your email address",
    		"email": "Provided email is not valid",
        },
        "new_password": {
    		"required": "Must provide password",
    		"min": "Password should be at-least 8 characters",
        },
        "confirm_password": {
    		"match": "Password should match New Password",
        },
    }
    return rules, messages
}
```
Here you can see we are returning `rules` of validation and custom `messages` for each validation rule of each `struct` field.

#### Running Validation
To use this `struct` for validation, you can use the following way.

```go
package controllers

import (
	"encoding/json"
	"github.com/go4all/validation"
	"log"
	"requests"
)

import "net/http"

type UserController struct{}

func (uc UserController) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	//...
	request := requests.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	vErrors, err := validation.Run(request)
	// Check if validation process ran successfully
	if err != nil {
		log.Fatal(err)
	} 
	// Check if we are having any validation errors 
	if len(vErrors) > 0 { 
		// respond with 442 status code and error messages 
		return
	} 
	// Continue with request data. 
	//...
}
```

```go
package main

import (
	"log"
	"net/http"
)

//...
func main() {
	mux := http.NewServeMux()
	userController := controllers.UserController{}
	mux.HandleFunc("/user/signup", userController.SignUpHandler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
```

#### Existing Validation Rules
Following table show existing validation rules. We will keep adding more rules.

| Name     | Description                                                         | Example                                    |
|----------|---------------------------------------------------------------------|--------------------------------------------|
| alphaNum | Field value should only contain letters and numbers and underscores | "username": {"alphaNum"}                   |
| email    | Field value is a valid email address                                | "email": {"email"}                         |
| match    | Match current field value matches with another field value          | "new_password": {"match:confirm_password"} |
| max      | Field should have max value                                         | "age": {"max:65"}                          |
| min      | Field should have min value                                         | "skills": {"min:5"}                        |
| regex    | Field value should match according to regular expression            | "level": {"regex:^(A\|B\|C)$"}             |
| required | Field value should exist                                            | "first_name": {"required"}                 |

#### Creating Custom Validation
You can also create custom validation rules by using one of two ways.

- Implement `Rule` interface
- Create new `func` of type `RuleCheck`

Custom rule example

```go
package example

import (
	"github.com/go4all/validation/types"
)

// Custom rule by implementing Rule interface
type PhoneRule struct{}

func (rule PhoneRule) GetError(kind reflect.Kind, field string, args []string) error {
	return "Default error message"
}

func (rule PhoneRule) Check(config types.RuleConfig) error {
	// do your check and return nil or error
}

// Custom rule by creating func of type `RuleCheck`
func IpRule(config types.RuleCheck) error {
	// do your check and return nil or error
}
```
Register custom rule before using it.
```go
package example

import (
	"github.com/go4all/validation"
)

func init() {
	// adding rule implementing `Rule` interface
	validation.AddRuleCheck("phone", PhoneRule{}.Check)
	// adding func of type `types.RuleCheck`
	validation.AddRuleCheck("ip", IpRule)
}
```

