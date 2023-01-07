package main

import (
	"app/app"
	devclip_service "app/pkg/devclip"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func main() {
	if err := run(); err != nil {
		logger.Error("startup error", zap.Error(err))
		os.Exit(1)
	}
}

func run() error {
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("error flushing log buffer", zap.Error(err))
		}
	}()

	router := gin.Default()
	router.Use(cors.Default())

	devclipService := devclip_service.NewDevclipService()

	server := app.NewServer(router, devclipService)

	// start the server
	err := server.Run()
	if err != nil {
		return err
	}

	return nil
}
