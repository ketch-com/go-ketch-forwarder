package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteResponseBody struct {
	Status                      RequestStatus       `json:"status"`
	Reason                      RequestStatusReason `json:"reason"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp"`
	RedirectURL                 string              `json:"redirectUrl,omitempty"`
	RequestID                   string              `json:"requestID,omitempty"`
	Documents                   []*Callback         `json:"documents,omitempty"`
	DocumentData                any                 `json:"documentData,omitempty"`
	Claims                      map[string]any      `json:"claims,omitempty"`
	Subject                     *DataSubject        `json:"subject,omitempty"`
	Identities                  []*Identity         `json:"identities,omitempty"`
	Messages                    []*Message          `json:"messages,omitempty"`
}

func (r *DeleteResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
