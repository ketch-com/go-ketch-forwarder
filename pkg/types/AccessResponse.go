package types

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AccessResponseBody struct {
	Status                      RequestStatus       `json:"status,omitempty"`
	Reason                      RequestStatusReason `json:"reason,omitempty"`
	ExpectedCompletionTimestamp int64               `json:"expectedCompletionTimestamp,omitempty"`
	RedirectURL                 string              `json:"redirectUrl,omitempty"`
	RequestID                   string              `json:"requestID,omitempty"`
	Results                     []*Callback         `json:"results,omitempty"`
	Documents                   []*Callback         `json:"documents,omitempty"`
	ResultData                  any                 `json:"resultData,omitempty"`
	DocumentData                any                 `json:"documentData,omitempty"`
	Claims                      map[string]any      `json:"claims,omitempty"`
	Subject                     *DataSubject        `json:"subject,omitempty"`
	Identities                  []*Identity         `json:"identities,omitempty"`
	Messages                    []*Message          `json:"messages,omitempty"`
}

func (r *AccessResponseBody) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, r,
		validation.Field(&r.Status, validation.Required, validation.In(RequestStatuses...)),
		validation.Field(&r.Reason, validation.When(len(r.Reason) > 0, validation.In(RequestStatusReasons...))),
		validation.Field(&r.ExpectedCompletionTimestamp, validation.Required),
	)
}
