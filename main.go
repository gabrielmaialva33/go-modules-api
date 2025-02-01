package main

import (
	"go-modules-api/config"
	"go-modules-api/internal/server/http"
	"go-modules-api/utils"

	"go.uber.org/zap"
)

func main() {

	utils.InitLogger()
	log := utils.Logger

	logger := log.Named("main")
	logger.Info("Starting the application")

	config.Load(log)

	logger.Info("server started", zap.Int("port", config.Env.AppPort))
	logger.Sugar().Infof("server is running at %s", config.Env.AppHost)

	s := http.NewServer(log)
	s.Start()
}
