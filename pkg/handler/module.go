package handler

import "go.uber.org/fx"

var Module = fx.Module("handlers",
	fx.Provide(
		NewSampleAccessRequestHandler,
		NewSampleCorrectionRequestHandler,
		NewSampleDeleteRequestHandler,
		NewSampleRestrictProcessingRequestHandler,
	),
)
