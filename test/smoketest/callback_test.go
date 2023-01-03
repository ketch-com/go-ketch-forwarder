//go:build !unit && !integration && smoke

package smoketest

import (
	"context"
	"encoding/json"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"log"
	"time"
)

func (suite *SmokeTestSuite) TestAccessCallback() {
	ctx := context.Background()
	//client := suite.params.Callbacks.Provide(ctx, suite.params.Config.CallbackURL)
	err := suite.params.Callbacks.SendAccessStatusEvent(ctx, &types.AccessStatusEvent{
		ApiVersion: "",
		Kind:       "",
		Metadata: &types.Metadata{
			UID:    "",
			Tenant: "",
		},
		Event: &types.AccessResponseBody{
			Status:                      "",
			Reason:                      "",
			ExpectedCompletionTimestamp: 0,
			RedirectURL:                 "",
			RequestID:                   "",
			Results:                     nil,
			Documents:                   nil,
			ResultData:                  nil,
			DocumentData:                nil,
			Claims:                      nil,
			Subject:                     nil,
			Identities:                  nil,
			Messages:                    nil,
		},
	}, "https://example.com/callback", make(map[string]string))
	if err.Error() == "Not Found" {
		err = nil
	}
	suite.Require().NoError(err)
}

func (suite *SmokeTestSuite) TestAccessCallbackFromInput() {
	ctx := context.Background()

	request := new(types.Request)
	err := json.Unmarshal([]byte(accessRequest), &request)
	if err != nil {
		log.Println("Error")
	}

	body := new(types.AccessRequestBody)
	if err := json.Unmarshal(request.Request, &body); err != nil {
		log.Println("Error")
	}

	accessRequest1 := &types.AccessStatusEvent{
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
		log.Println(cbk.URL, cbk.Headers)
		err := suite.params.Callbacks.SendAccessStatusEvent(ctx, accessRequest1, cbk.URL, cbk.Headers)
		if err.Error() == "Not Found" {
			err = nil
		}
		suite.Require().NoError(err)
	}
}
