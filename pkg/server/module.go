package server

import (
	"github.com/ketch-com/go-ketch-forwarder/pkg/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("server",
	handler.Module,

	fx.Provide(
		NewConfig,
		NewListener,
		NewServer,
		NewHandler,
	),

	fx.Invoke(
		Lifecycle,
	),
)
