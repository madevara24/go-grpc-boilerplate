package auth

import (
	"go-grpc-boilerplate/internal/app"
	authpb "go-grpc-boilerplate/proto/auth"
)

type AuthHandler struct {
	container *app.Container
	authpb.UnimplementedAuthHandlerServer
}

func NewAuthHandler(container *app.Container) *AuthHandler {
	return &AuthHandler{
		container: container,
	}
}
