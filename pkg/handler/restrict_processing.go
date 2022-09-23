package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"time"
)

type SampleRestrictProcessingRequestHandler struct{}

func NewSampleRestrictProcessingRequestHandler() types.RestrictProcessingHandler {
	return &SampleRestrictProcessingRequestHandler{}
}

func (h *SampleRestrictProcessingRequestHandler) RestrictProcessing(ctx context.Context, request *types.RestrictProcessingRequestBody) (*types.RestrictProcessingResponseBody, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	resp := &types.RestrictProcessingResponseBody{
		Status:                      types.PendingRequestStatus,
		ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
	}

	return resp, nil
}
