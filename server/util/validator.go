package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Explanations are a map keyed by the field, values are Explanations
type Explanations map[string]Explanation

// Explanation is a map keyed by the rule, value is the error message
type Explanation map[string]string

// ValidateRequest returns a GRPC error response
func ValidateRequest(in interface{}, errorMap Explanations) error {

	errs := validator.New().Struct(in)
	if errs == nil {
		return nil
	}

	errors := errs.(validator.ValidationErrors)

	s := status.New(codes.InvalidArgument, "Invalid argument values")

	violations := []*errdetails.BadRequest_FieldViolation{}
	for _, err := range errors {
		desc := fmt.Sprintf("Validation rule \"%s\" failed", err.Tag())
		if explanations, ok := errorMap[err.Field()]; ok {
			if explanation, okDepth := explanations[err.Tag()]; okDepth {
				desc = explanation
			}
		}

		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       err.Field(),
			Description: desc,
		})
	}

	sWD, _ := s.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: violations,
		},
	)

	proto := sWD.Proto()
	return status.ErrorProto(proto)
}
