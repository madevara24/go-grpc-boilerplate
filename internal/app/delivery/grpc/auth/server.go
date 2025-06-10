package auth

import (
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/proto/auth"
)

type AuthServer struct {
	container *app.Container
	auth.UnimplementedAuthServiceServer
}

func NewAuthServer(container *app.Container) *AuthServer {
	return &AuthServer{
		container: container,
	}
}
