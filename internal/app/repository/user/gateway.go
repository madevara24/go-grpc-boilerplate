package user

import (
	"context"
	"go-grpc-boilerplate/internal/app/entity"
	"go-grpc-boilerplate/internal/pkg/datasource"
)

type IRepo interface {
	Create(ctx context.Context, req entity.User) error
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}

type repo struct {
	datasource *datasource.DataSource
}

func NewRepo(datasource *datasource.DataSource) IRepo {
	return &repo{
		datasource: datasource,
	}
}
