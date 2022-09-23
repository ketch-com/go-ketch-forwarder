package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ketch-com/go-ketch-forwarder/pkg/metadata"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type HandlerParams struct {
	fx.In

	Config             Config
	Access             types.AccessHandler             `optional:"true"`
	Delete             types.DeleteHandler             `optional:"true"`
	Correction         types.CorrectionHandler         `optional:"true"`
	RestrictProcessing types.RestrictProcessingHandler `optional:"true"`
}

type kindHandler func(ctx context.Context, w http.ResponseWriter, request *types.Request)

type Handler struct {
	params   HandlerParams
	handlers map[types.Kind]kindHandler
}

func NewHandler(params HandlerParams) http.Handler {
	handler := &Handler{
		params:   params,
		handlers: make(map[types.Kind]kindHandler),
	}

	if params.Access != nil {
		handler.handlers[types.AccessRequestKind] = makeHandler[types.AccessRequestBody, types.AccessResponseBody](types.AccessResponseKind, params.Access.Access)
	}

	if params.Delete != nil {
		handler.handlers[types.DeleteRequestKind] = makeHandler[types.DeleteRequestBody, types.DeleteResponseBody](types.DeleteResponseKind, params.Delete.Delete)
	}

	if params.Correction != nil {
		handler.handlers[types.CorrectionRequestKind] = makeHandler[types.CorrectionRequestBody, types.CorrectionResponseBody](types.CorrectionResponseKind, params.Correction.Correction)
	}

	if params.RestrictProcessing != nil {
		handler.handlers[types.RestrictProcessingRequestKind] = makeHandler[types.RestrictProcessingRequestBody, types.RestrictProcessingResponseBody](types.RestrictProcessingResponseKind, params.RestrictProcessing.RestrictProcessing)
	}

	mux := chi.NewMux()
	if len(params.Config.Username) > 0 && len(params.Config.Password) > 0 {
		mux.Use(BasicAuth(params.Config.Username, params.Config.Password))
	} else {
		log.Println("⚠️ Starting server without authentication. Not recommended for production use.")
	}

	mux.Post("/ketch/events", handler.ServeHTTP)
	return mux
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		WriteError(w, nil, http.StatusBadRequest, "invalid", fmt.Sprintf("expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type")))
		return
	}

	if !CanAccept(r, "application/json") {
		WriteError(w, nil, http.StatusBadRequest, "invalid", fmt.Sprintf("expected Accept to include 'application/json', got '%s'", r.Header.Get("Accept")))
		return
	}

	request := new(types.Request)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteError(w, nil, http.StatusBadRequest, "invalid", err.Error())
		return
	}

	if err := validation.ValidateWithContext(ctx, request); err != nil {
		WriteError(w, request.Metadata, http.StatusBadRequest, "invalid", err.Error())
		return
	}

	ctx = metadata.WithRequest(ctx, request)

	if handler, ok := h.handlers[request.Kind]; ok {
		handler(ctx, w, request)
		return
	}

	WriteError(w, request.Metadata, http.StatusBadRequest, "invalid", fmt.Sprintf("invalid request kind '%s'", request.Kind))
}

func makeHandler[Request any, Response any](kind types.Kind, handler func(context.Context, *Request) (*Response, error)) kindHandler {
	return func(ctx context.Context, w http.ResponseWriter, request *types.Request) {
		body := new(Request)

		if err := json.Unmarshal(request.Request, &body); err != nil {
			WriteError(w, request.Metadata, http.StatusBadRequest, "invalid", err.Error())
			return
		}

		if err := validation.ValidateWithContext(ctx, body); err != nil {
			WriteError(w, request.Metadata, http.StatusBadRequest, "invalid", err.Error())
			return
		}

		resp, err := handler(ctx, body)
		if err != nil {
			WriteError(w, request.Metadata, http.StatusInternalServerError, "internal", err.Error())
			return
		}

		out := &types.Response{
			ApiVersion: request.ApiVersion,
			Kind:       kind,
			Metadata:   request.Metadata,
		}

		out.Response, err = json.Marshal(resp)
		if err != nil {
			WriteError(w, request.Metadata, http.StatusInternalServerError, "internal", err.Error())
			return
		}

		if err = validation.ValidateWithContext(ctx, out); err != nil {
			WriteError(w, request.Metadata, http.StatusInternalServerError, "internal", err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(out); err != nil {
			WriteError(w, request.Metadata, http.StatusInternalServerError, "internal", err.Error())
			return
		}
	}
}
