package cmd

import (
	"context"
	"go-grpc-boilerplate/config"
	"go-grpc-boilerplate/internal/app"
	"go-grpc-boilerplate/internal/app/delivery/rest"
	"go-grpc-boilerplate/internal/pkg/datasource"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/madevara24/go-common/server"
	ginCommon "github.com/madevara24/go-common/server/gin"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srvConfig := server.Config{
		Address:        config.Get().Port,
		Env:            config.Get().ENV,
		ReadTimeout:    time.Duration(config.Get().ServerReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Get().ServerWriteTimeout) * time.Second,
		AllowedOrigins: config.Get().AllowedOrigins,
	}

	ginHttpServer, err := ginCommon.NewGinHttpServer(srvConfig)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(ginHttpServer.GetRouter(), srvConfig)

	datasource := datasource.NewDataSource()

	container := app.NewContainer(datasource)

	router := rest.NewRouter(ctx, ginHttpServer.GetRouter(), container)

	router.RegisterRouter()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	srv.StartServer()

	<-quit
}
