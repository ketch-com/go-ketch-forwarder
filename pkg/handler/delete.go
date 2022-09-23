package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"time"
)

type SampleDeleteRequestHandler struct{}

func NewSampleDeleteRequestHandler() types.DeleteHandler {
	return &SampleDeleteRequestHandler{}
}

func (h *SampleDeleteRequestHandler) Delete(ctx context.Context, request *types.DeleteRequestBody) (*types.DeleteResponseBody, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	resp := &types.DeleteResponseBody{
		Status:                      types.PendingRequestStatus,
		Reason:                      types.OtherRequestStatusReason,
		ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
	}

	return resp, nil
}
