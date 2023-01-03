package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessRequestBody struct {
	Controller         string         `json:"controller,omitempty"`
	Property           string         `json:"property,omitempty"`
	Environment        string         `json:"environment,omitempty"`
	Regulation         string         `json:"regulation,omitempty"`
	Jurisdiction       string         `json:"jurisdiction,omitempty"`
	Identities         []*Identity    `json:"identities,omitempty"`
	Callbacks          []*Callback    `json:"callbacks,omitempty"`
	Subject            *DataSubject   `json:"subject,omitempty"`
	Claims             map[string]any `json:"claims,omitempty"`
	SubmittedTimestamp int64          `json:"submittedTimestamp,omitempty"`
	DueTimestamp       int64          `json:"dueTimestamp,omitempty"`
}

func (r *AccessRequestBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Property, validation.Required),
		validation.Field(&r.Environment, validation.Required),
		validation.Field(&r.Regulation, validation.Required),
		validation.Field(&r.Jurisdiction, validation.Required),
		validation.Field(&r.Identities, validation.Required),
		validation.Field(&r.Subject, validation.Required),
	)
}
