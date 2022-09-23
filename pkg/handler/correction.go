package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"time"
)

type SampleCorrectionRequestHandler struct{}

func NewSampleCorrectionRequestHandler() types.CorrectionHandler {
	return &SampleCorrectionRequestHandler{}
}

func (h *SampleCorrectionRequestHandler) Correction(ctx context.Context, request *types.CorrectionRequestBody) (*types.CorrectionResponseBody, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	resp := &types.CorrectionResponseBody{
		Status:                      types.PendingRequestStatus,
		Reason:                      types.OtherRequestStatusReason,
		ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
	}

	return resp, nil
}
