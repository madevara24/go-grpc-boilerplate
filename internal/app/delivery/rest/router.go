package rest

import (
	"context"
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/internal/app/delivery/rest/auth"
	"go-grpc-boilerplate/internal/app/delivery/rest/healthcheck"
	"go-grpc-boilerplate/internal/app/delivery/rest/middleware"
	"go-grpc-boilerplate/internal/app/delivery/rest/user"

	"github.com/gin-gonic/gin"
)

type Router struct {
	context   context.Context // NOTE: For middlewares
	router    *gin.Engine
	container *app.Container
}

func NewRouter(ctx context.Context, router *gin.Engine, container *app.Container) *Router {
	return &Router{
		context:   ctx,
		router:    router,
		container: container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.Use(gin.Recovery())

	v1 := h.router.Group("/v1")

	// PING
	v1.GET("/health", healthcheck.HealthCheck(h.container.HealthCheckInport))

	// AUTH
	authRouter := v1.Group("/auth")
	authRouter.POST("/login", auth.Login(h.container.AuthLoginInport))
	authRouter.POST("/refresh", auth.Refresh(h.container.AuthRefreshInport))

	// USER
	userRouter := v1.Group("/user")
	userRouter.POST("/register", user.Register(h.container.UserRegisterInport))

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Add protected routes here
	}
}
