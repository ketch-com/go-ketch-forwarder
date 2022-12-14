package types

import (
	"context"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Request struct {
	ApiVersion string          `json:"apiVersion,omitempty"`
	Kind       Kind            `json:"kind,omitempty"`
	Metadata   *Metadata       `json:"metadata,omitempty"`
	Request    json.RawMessage `json:"request,omitempty"`
}

func (r *Request) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.ApiVersion, validation.Required, validation.In(ApiVersion)),
		validation.Field(&r.Kind, validation.Required, validation.In(Kinds...)),
		validation.Field(&r.Metadata, validation.Required),
		validation.Field(&r.Request, validation.Required),
	)
}
