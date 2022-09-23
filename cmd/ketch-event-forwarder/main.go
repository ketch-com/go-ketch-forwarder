package main

import (
	"context"
	"github.com/ketch-com/go-ketch-forwarder/pkg/server"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		server.Module,
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalln(err)
	}

	defer app.Stop(ctx)

	<-app.Done()

	if err := app.Err(); err != nil {
		log.Fatalln(err)
	}
}
