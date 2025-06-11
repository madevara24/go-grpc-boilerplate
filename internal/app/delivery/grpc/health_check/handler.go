package health_check

import (
	"go-grpc-boilerplate/internal/app"
	healthpb "go-grpc-boilerplate/proto/health_check"
)

type HealthCheckHandler struct {
	container *app.Container
	healthpb.UnimplementedHealthCheckHandlerServer
}

func NewHealthCheckHandler(container *app.Container) *HealthCheckHandler {
	return &HealthCheckHandler{
		container: container,
	}
}
