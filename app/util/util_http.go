package util

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"sw-planet-service-go/dto"
)

func ParseValidationFields(err error) []dto.ApiError {
	var ve validator.ValidationErrors
	errors.As(err, &ve)
	out := make([]dto.ApiError, len(ve))
	for i, fe := range ve {
		out[i] = dto.ApiError{Field: fe.Field(), Msg: "This field is required"}
	}
	return out
}
