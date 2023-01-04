package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"go.ketch.com/lib/orlop/v2/errors"
	"net/http"
	"time"
)

type CallbackClient struct {
	client http.Client
}

func NewCallbackClient() *CallbackClient {
	client := http.Client{Timeout: 2 * time.Second}
	return &CallbackClient{
		client: client,
	}
}

func (c *CallbackClient) SendAccessStatusEvent(ctx context.Context, request *types.AccessStatusEvent, url string, headers map[string]string) error {
	return c.send(ctx, request, url, headers)
}

func (c *CallbackClient) SendDeleteStatusEvent(ctx context.Context, request *types.DeleteStatusEvent, url string, headers map[string]string) error {
	return c.send(ctx, request, url, headers)
}

func (c *CallbackClient) SendCorrectionStatusEvent(ctx context.Context, request *types.CorrectionStatusEvent, url string, headers map[string]string) error {
	return c.send(ctx, request, url, headers)
}

func (c *CallbackClient) SendRestrictProcessingStatusEvent(ctx context.Context, request *types.RestrictProcessingStatusEvent, url string, headers map[string]string) error {
	return c.send(ctx, request, url, headers)
}

func (c *CallbackClient) send(ctx context.Context, request any, url string, headers map[string]string) error {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(request); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > 399 {
		errResp := &types.Error{}

		if resp.Header.Get("Content-Type") == "application/json" {
			err = json.NewDecoder(resp.Body).Decode(&errResp)
			if err != nil {
				return err
			}

			return errors.WithStatusCode(errors.New(errResp.Error.Message), errResp.Error.Code)
		}

		return errors.WithStatusCode(nil, resp.StatusCode)
	}

	defer resp.Body.Close()

	return nil
}
