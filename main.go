package main

import (
	"os"

	"go-modules-api/cmd"
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

	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	logger.Info("server started", zap.Int("port", config.Env.AppPort))
	logger.Sugar().Infof("server is running at %s", config.Env.AppHost)

	s := http.NewServer(log)
	s.Start()
}
