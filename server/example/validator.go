package example

import (
	va "github.com/go-playground/validator/v10"
	p "github.com/journeyai/grpc-validation/protocols/example"
	"github.com/journeyai/grpc-validation/server/util"
)

// Validator validates requests
type Validator struct {
	validate va.Validate
}

func newValidator() Validator {
	return Validator{
		validate: *va.New(),
	}
}

// ValidateCreateWidget validates CreateWidget requests
func (v Validator) ValidateCreateWidget(in p.CreateWidgetRequest) error {

	explanations := util.Explanations{
		"Name": util.Explanation{
			"required": "A name is required",
		},
		"Age": util.Explanation{
			"required": "An age is required for age verification",
			"gte":      "Age must be at least 21 years old",
		},
		"Email": util.Explanation{
			"required": "An email address is required",
			"email":    "A valid email address is required",
		},
		"PhoneNumber": util.Explanation{
			"e164": "A phone number in e164 format is required (+13035551999)",
		},
	}

	errs := v.validate.Struct(in)
	if errs == nil {
		return nil
	}

	errors := errs.(va.ValidationErrors)
	return util.HandleErrorResponse(errors, explanations)
}

// ValidateGetWidget validates GetWidget requests
func (v Validator) ValidateGetWidget(in p.GetWidgetRequest) error {

	explanations := util.Explanations{
		"Id": util.Explanation{
			"required": "An id reference for a user is required",
			"uuid4":    "A UUIDv4 format is required for the id",
		},
	}

	errs := v.validate.Struct(in)
	if errs == nil {
		return nil
	}

	errors := errs.(va.ValidationErrors)
	return util.HandleErrorResponse(errors, explanations)
}
