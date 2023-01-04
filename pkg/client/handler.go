package client

import (
	"context"
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"net/http"
	"time"
)

type Handler struct {
	CallbackClient CallbackClient
}

func NewCallbackHandler(callbackClient CallbackClient) *Handler {
	return &Handler{CallbackClient: callbackClient}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		WriteError(w, nil, http.StatusBadRequest, "invalid", fmt.Sprintf("expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type")))
		return
	}

	if r.Header.Get("Accept") != "application/json" {
		WriteError(w, nil, http.StatusBadRequest, "invalid", fmt.Sprintf("expected Accept to include 'application/json', got '%s'", r.Header.Get("Accept")))
		return
	}

	req := new(types.Request)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, nil, http.StatusBadRequest, "invalid", err.Error())
		return
	}

	if err := validation.ValidateWithContext(ctx, req); err != nil {
		WriteError(w, req.Metadata, http.StatusBadRequest, "invalid", err.Error())
		return
	}

	err := h.HandleEvent(ctx, req)
	if err != nil {
		WriteError(w, req.Metadata, http.StatusBadRequest, "invalid", err.Error())
		return
	}

	out := &types.Response{
		ApiVersion: req.ApiVersion,
		Kind:       req.Kind,
		Metadata:   req.Metadata,
	}

	//out.Response, err = json.Marshal(resp)
	//if err != nil {
	//	WriteError(w, req.Metadata, http.StatusInternalServerError, "internal", err.Error())
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err = json.NewEncoder(w).Encode(out); err != nil {
		WriteError(w, req.Metadata, http.StatusInternalServerError, "internal", err.Error())
		return
	}

}

func (h *Handler) HandleEvent(ctx context.Context, request *types.Request) error {

	if request.Kind == types.AccessRequestKind {
		err := h.processAccessStatusEvent(ctx, request)
		if err != nil {
			return err
		}
	} else if request.Kind == types.DeleteRequestKind {
		err := h.processDeleteStatusEvent(ctx, request)
		if err != nil {
			return err
		}
	} else if request.Kind == types.CorrectionRequestKind {
		err := h.processCorrectionStatusEvent(ctx, request)
		if err != nil {
			return err
		}
	} else if request.Kind == types.RestrictProcessingRequestKind {
		err := h.processRestrictStatusEvent(ctx, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) processAccessStatusEvent(ctx context.Context, request *types.Request) error {
	body := new(types.AccessRequestBody)
	if err := json.Unmarshal(request.Request, &body); err != nil {
		return err
	}

	accessEvent := &types.AccessStatusEvent{
		ApiVersion: request.ApiVersion,
		Kind:       request.Kind,
		Metadata:   request.Metadata,
		Event: &types.AccessResponseBody{
			Status:                      types.InProgressRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
			RedirectURL:                 "",
			RequestID:                   "",
			Results:                     nil,
			Documents:                   nil,
			ResultData:                  nil,
			DocumentData:                nil,
			Claims:                      body.Claims,
			Subject:                     body.Subject,
			Identities:                  body.Identities,
			Messages:                    nil,
		},
	}
	for _, cbk := range body.Callbacks {
		err := h.CallbackClient.SendAccessStatusEvent(ctx, accessEvent, cbk.URL, cbk.Headers)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) processDeleteStatusEvent(ctx context.Context, request *types.Request) error {
	body := new(types.DeleteRequestBody)
	if err := json.Unmarshal(request.Request, &body); err != nil {
		return err
	}

	deleteEvent := &types.DeleteStatusEvent{
		ApiVersion: request.ApiVersion,
		Kind:       request.Kind,
		Metadata:   request.Metadata,
		Event: &types.DeleteResponseBody{
			Status:                      types.InProgressRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
			RedirectURL:                 "",
			RequestID:                   "",
			Documents:                   nil,
			DocumentData:                nil,
			Claims:                      body.Claims,
			Subject:                     body.Subject,
			Identities:                  body.Identities,
			Messages:                    nil,
		},
	}

	for _, cbk := range body.Callbacks {
		err := h.CallbackClient.SendDeleteStatusEvent(ctx, deleteEvent, cbk.URL, cbk.Headers)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) processCorrectionStatusEvent(ctx context.Context, request *types.Request) error {
	body := new(types.CorrectionRequestBody)
	if err := json.Unmarshal(request.Request, &body); err != nil {
		return err
	}

	correctionEvent := &types.CorrectionStatusEvent{
		ApiVersion: request.ApiVersion,
		Kind:       request.Kind,
		Metadata:   request.Metadata,
		Event: &types.CorrectionResponseBody{
			Status:                      types.InProgressRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
			RedirectURL:                 "",
			RequestID:                   "",
			Documents:                   nil,
			DocumentData:                nil,
			Claims:                      body.Claims,
			Subject:                     body.Subject,
			Identities:                  body.Identities,
			Messages:                    nil,
		},
	}

	for _, cbk := range body.Callbacks {
		err := h.CallbackClient.SendCorrectionStatusEvent(ctx, correctionEvent, cbk.URL, cbk.Headers)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) processRestrictStatusEvent(ctx context.Context, request *types.Request) error {
	body := new(types.RestrictProcessingRequestBody)
	if err := json.Unmarshal(request.Request, &body); err != nil {
		return err
	}

	restrictEvent := &types.RestrictProcessingStatusEvent{
		ApiVersion: request.ApiVersion,
		Kind:       request.Kind,
		Metadata:   request.Metadata,
		Event: &types.RestrictProcessingResponseBody{
			Status:                      types.InProgressRequestStatus,
			Reason:                      types.OtherRequestStatusReason,
			ExpectedCompletionTimestamp: time.Now().Unix(),
			RedirectURL:                 "",
			RequestID:                   "",
			Documents:                   nil,
			DocumentData:                nil,
			Claims:                      body.Claims,
			Subject:                     body.Subject,
			Identities:                  body.Identities,
			Messages:                    nil,
		},
	}

	for _, cbk := range body.Callbacks {
		err := h.CallbackClient.SendRestrictProcessingStatusEvent(ctx, restrictEvent, cbk.URL, cbk.Headers)
		if err != nil {
			return err
		}
	}

	return nil
}
