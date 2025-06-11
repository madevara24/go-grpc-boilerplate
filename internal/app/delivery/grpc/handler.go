package grpc

import (
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/internal/app/delivery/grpc/auth"
	"go-grpc-boilerplate/internal/app/delivery/grpc/health_check"
	authpb "go-grpc-boilerplate/proto/auth"
	healthpb "go-grpc-boilerplate/proto/health_check"

	"google.golang.org/grpc"
)

type Handler struct {
	GrpcServer *grpc.Server

	AuthServer   *auth.AuthHandler
	HealthServer *health_check.HealthCheckHandler
}

func NewHandler(server *grpc.Server, container *app.Container) *Handler {
	return &Handler{
		GrpcServer:   server,
		AuthServer:   auth.NewAuthHandler(container),
		HealthServer: health_check.NewHealthCheckHandler(container),
	}
}

func (s *Handler) RegisterHandler() {

	authpb.RegisterAuthHandlerServer(s.GrpcServer, s.AuthServer)
	healthpb.RegisterHealthCheckHandlerServer(s.GrpcServer, s.HealthServer)
}
