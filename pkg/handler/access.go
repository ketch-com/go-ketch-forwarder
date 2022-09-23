package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"time"
)

type SampleAccessRequestHandler struct{}

func NewSampleAccessRequestHandler() types.AccessHandler {
	return &SampleAccessRequestHandler{}
}

func (h *SampleAccessRequestHandler) Access(ctx context.Context, request *types.AccessRequestBody) (*types.AccessResponseBody, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	resp := &types.AccessResponseBody{
		Status:                      types.PendingRequestStatus,
		Reason:                      types.OtherRequestStatusReason,
		ExpectedCompletionTimestamp: time.Now().Add(45 * 24 * time.Hour).Unix(),
	}

	return resp, nil
}
