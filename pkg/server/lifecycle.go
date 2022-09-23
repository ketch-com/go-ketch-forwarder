package server

import (
	"context"
	"go.uber.org/fx"
	"log"
	"net"
	"net/http"
)

// Lifecycle runs the HTTP Server
func Lifecycle(lifecycle fx.Lifecycle, listener net.Listener, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return start(listener, server) },
		OnStop:  func(ctx context.Context) error { return stop(ctx, server) },
	})
}

func start(listener net.Listener, server *http.Server) error {
	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			log.Fatalln("failed to start server")
		} else if err == http.ErrServerClosed {
			log.Println("server stopped")
		}
	}()

	return nil
}

func stop(ctx context.Context, server *http.Server) error {
	log.Println("stopping server")
	return server.Shutdown(ctx)
}
