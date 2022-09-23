package types

import "context"

type AccessHandler interface {
	Access(ctx context.Context, request *AccessRequestBody) (*AccessResponseBody, error)
}

type DeleteHandler interface {
	Delete(ctx context.Context, request *DeleteRequestBody) (*DeleteResponseBody, error)
}

type CorrectionHandler interface {
	Correction(ctx context.Context, request *CorrectionRequestBody) (*CorrectionResponseBody, error)
}

type RestrictProcessingHandler interface {
	RestrictProcessing(ctx context.Context, request *RestrictProcessingRequestBody) (*RestrictProcessingResponseBody, error)
}
