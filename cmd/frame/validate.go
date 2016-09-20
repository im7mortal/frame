package main

import (
	"github.com/asaskevich/govalidator"
)


type resp struct {
	StatusCode int `json:"statusCode"`
	Error string `json:"error"`
	Message string `json:"message"`
	Validation map[string]interface{}
}

func validate(o interface{}) (bool, *resp)  {

	valid, err := govalidator.ValidateStruct(o)
	if valid {
		return true, nil
	}
	if err != nil {
		if errors, ok := err.(govalidator.Errors); ok {
			res := resp{}
			res.StatusCode = 400
			res.Error = "Bad Request"
			res.Message = errors[0].Error()
			res.Validation = map[string]interface{}{
				"source": "payload",
				"keys": map[string]interface{}{
					"source": "email",
				},
			}
			return false, &res
		}
		EXCEPTION(err)
	}
	return false, nil
}

/*

{
  "statusCode": 400,
  "error": "Bad Request",
  "message": "child \"email\" fails because [\"email\" is required]",
  "validation": {
    "source": "payload",
    "keys": [
      "email"
    ]
  }
}

*/