package health_check

import (
	"context"
	healthpb "go-grpc-boilerplate/proto/health_check"
)

func (s *HealthCheckHandler) HealthCheck(ctx context.Context, in *healthpb.HealthRequest) (*healthpb.HealthResponse, error) {

	_ = s.container.HealthCheckInport.Execute(ctx)

	return &healthpb.HealthResponse{}, nil
}
