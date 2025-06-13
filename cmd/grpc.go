package cmd

import (
	"context"
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/internal/app/delivery/grpc"
	"go-grpc-boilerplate/internal/pkg/datasource"
	"os"
	"os/signal"
	"syscall"

	grpc_server "go-grpc-boilerplate/internal/pkg/grpc-server"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Runs the grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		runGrpc()
	},
}

func runGrpc() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcServer := grpc_server.NewGRPCServer()

	datasource := datasource.NewDataSource()

	container := app.NewContainer(datasource)

	handler := grpc.NewHandler(ctx, grpcServer.GetGRPCServer(), container)

	handler.RegisterHandler()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	grpcServer.StartServer()

	<-quit
}
