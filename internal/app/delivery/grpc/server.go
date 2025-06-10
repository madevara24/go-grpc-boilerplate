package grpc

import (
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/internal/app/delivery/grpc/auth"
	authpb "go-grpc-boilerplate/proto/auth"

	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *grpc.Server

	AuthServer *auth.AuthServer
}

func NewServer(server *grpc.Server, container *app.Container) *Server {
	return &Server{
		GrpcServer: server,
		AuthServer: auth.NewAuthServer(container),
	}
}

func (s *Server) RegisterServer() {

	authpb.RegisterAuthServiceServer(s.GrpcServer, s.AuthServer)
}
