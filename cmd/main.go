package main

import (
	"app/app"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func main() {
	if err := run(); err != nil {
		logger.Error("this is the startup error", zap.Error(err))
		os.Exit(1)
	}
}

func run() error {
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("could not flush logger buffer", zap.Error(err))
		}
	}()

	router := gin.Default()
	router.Use(cors.Default())

	server := app.NewServer(router)

	// start the server
	err := server.Run()
	if err != nil {
		return err
	}

	return nil
}
