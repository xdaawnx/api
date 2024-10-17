package libraries

import (
	util "api/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

var validationMessages = map[string]string{
	"required":    "%s is required",
	"email":       "%s must be a valid email address",
	"url":         "%s must be a valid URL",
	"uuid":        "%s must be a valid UUID",
	"min":         "%s must be at least %s",
	"max":         "%s must not be greater than %s",
	"len":         "%s must be exactly %s characters long",
	"gte":         "%s must be greater than or equal to %s",
	"lte":         "%s must be less than or equal to %s",
	"gt":          "%s must be greater than %s",
	"lt":          "%s must be less than %s",
	"alpha":       "%s must contain only letters",
	"alphanum":    "%s must contain only letters and numbers",
	"numeric":     "%s must be a number",
	"ascii":       "%s must contain only ASCII characters",
	"printascii":  "%s must contain only printable ASCII characters",
	"datetime":    "%s must be a valid datetime",
	"json":        "%s must be a valid JSON",
	"startswith":  "%s must start with %s",
	"endswith":    "%s must end with %s",
	"contains":    "%s must contain %s",
	"not":         "%s must not be %s",
	"oneof":       "%s must be one of: %s",
	"divisibleby": "%s must be divisible by %s",
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.Validator.Struct(i); err != nil {
		validationErrors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			// Get the custom message format for the validation tag
			msgTemplate, exists := validationMessages[err.Tag()]
			if !exists {
				msgTemplate = "%s is invalid" // Default message for unknown tags
			}

			// Replace placeholders with field name and tag parameters
			var msg string
			switch err.Tag() {
			case "min", "max", "len", "gte", "lte", "gt", "lt", "oneof", "divisibleby", "not":
				msg = fmt.Sprintf(msgTemplate, err.Field(), err.Param())
			default:
				msg = fmt.Sprintf(msgTemplate, err.Field())
			}
			validationErrors[err.Field()] = msg
		}
		// Join validation errors into a single error message
		util.BadRequest.Validation = validationErrors
		return util.BadRequest
	}
	return nil
}
