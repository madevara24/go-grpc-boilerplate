package app

import (
	"go-grpc-boilerplate/internal/app/repository/user"
	"go-grpc-boilerplate/internal/app/usecase/auth/login"
	"go-grpc-boilerplate/internal/app/usecase/auth/refresh"
	"go-grpc-boilerplate/internal/app/usecase/healthcheck"
	"go-grpc-boilerplate/internal/app/usecase/user/register"
	"go-grpc-boilerplate/internal/pkg/datasource"
)

type Container struct {
	// PING
	HealthCheckInport healthcheck.Inport

	// USER
	UserRegisterInport register.Inport

	// AUTH
	AuthLoginInport   login.Inport
	AuthRefreshInport refresh.Inport
}

func NewContainer(datasource *datasource.DataSource) *Container {
	userRepo := user.NewRepo(datasource)
	return &Container{
		// PING
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),

		// USER
		UserRegisterInport: register.NewUsecase(userRepo),

		// AUTH
		AuthLoginInport:   login.NewUsecase(userRepo),
		AuthRefreshInport: refresh.NewUsecase(),
	}
}
