package example

import (
	context "context"
	fmt "fmt"

	"github.com/go-playground/validator/v10"
	"github.com/journeyai/grpc-validation/server/util"
)

// Server handles the Example protocol
type Server struct {
	validate validator.Validate
}

// NewServer creates a new server
func NewServer() (*Server, error) {

	s := Server{
		validate: *validator.New(),
	}

	return &s, nil
}

// CreateWidget handles the CreateWidget request
func (s Server) CreateWidget(ctx context.Context, in *CreateWidgetRequest) (*CreateWidgetReply, error) {

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

	if errs := s.validate.Struct(in); errs != nil {
		errors := errs.(validator.ValidationErrors)
		return nil, util.HandleErrorResponse(errors, explanations)
	}

	out := CreateWidgetReply{
		Id: "",
	}

	return &out, nil
}

// GetWidget handles the GetWidget request
func (s Server) GetWidget(ctx context.Context, in *GetWidgetRequest) (*GetWidgetReply, error) {

	return nil, fmt.Errorf("Theres an error (GetWidget)")
}
