package example

import (
	p "github.com/journeyai/grpc-validation/protocols/example"
	"github.com/journeyai/grpc-validation/server/util"
)

func validateCreateWidget(in p.CreateWidgetRequest) error {

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

	return util.ValidateRequest(in, explanations)
}

func validateGetWidget(in p.GetWidgetRequest) error {

	explanations := util.Explanations{
		"Id": util.Explanation{
			"required": "An id reference for a user is required",
			"uuid4":    "A UUIDv4 format is required for the id",
		},
	}

	return util.ValidateRequest(in, explanations)
}
