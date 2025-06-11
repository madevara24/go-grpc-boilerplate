package auth

import (
	"context"
	"go-grpc-boilerplate/internal/app/usecase/auth/refresh"
	authpb "go-grpc-boilerplate/proto/auth"
)

func (s *AuthHandler) Refresh(ctx context.Context, in *authpb.RefreshRequest) (*authpb.TokenResponse, error) {
	req := refresh.InportRequest{
		RefreshToken: in.RefreshToken,
	}

	token, err := s.container.AuthRefreshInport.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	return &authpb.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
