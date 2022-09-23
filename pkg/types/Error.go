package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ErrorBody struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *ErrorBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Code, validation.Required),
		validation.Field(&r.Status, validation.Required),
		validation.Field(&r.Message, validation.Required),
	)
}
