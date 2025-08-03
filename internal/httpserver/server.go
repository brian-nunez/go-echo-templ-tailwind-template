package httpserver

import (
	"context"

	v1 "github.com/brian-nunez/go-echo-starter-template/internal/handlers/v1"
)

type Server interface {
	Start(addr string) error
	Shutdown(ctx context.Context) error
}

func Bootstrap() Server {
	server := New().
		WithDefaultMiddleware().
		WithErrorHandler().
		WithRoutes(v1.RegisterRoutes).
		WithNotFound().
		Build()

	return server
}
