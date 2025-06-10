package auth

import (
	"context"
	"go-grpc-boilerplate/internal/app/usecase/auth/login"
	authpb "go-grpc-boilerplate/proto/auth"
)

func (s *AuthServer) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.TokenResponse, error) {
	req := login.InportRequest{
		Email:    in.Email,
		Password: in.Password,
	}

	token, err := s.container.AuthLoginInport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	return &authpb.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
