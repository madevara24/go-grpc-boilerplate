package grpc_server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type IGRPCServer interface {
	StartServer()
	GetGRPCServer() *grpc.Server
}

type grpcServer struct {
	grpcServer *grpc.Server
}

type Option func(*grpcServer)

func NewGRPCServer(opts ...Option) IGRPCServer {
	grpcServer := &grpcServer{
		grpcServer: grpc.NewServer(),
	}

	for _, opt := range opts {
		opt(grpcServer)
	}

	return grpcServer
}

func (g *grpcServer) GetGRPCServer() *grpc.Server {
	return g.grpcServer
}

func (g *grpcServer) StartServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := g.grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Println("gRPC server is running ...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")

	done := make(chan struct{})
	go func() {
		g.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server stopped gracefully")
	case <-time.After(time.Second * 5):
		log.Println("gRPC server shutdown timeout, forcing stop...")
		g.grpcServer.Stop()
		log.Println("gRPC server stopped.")
		os.Exit(1)
	}
}
