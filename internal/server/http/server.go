package http

import (
	"fmt"
	"go-modules-api/config"
	"go-modules-api/internal/server/container"
	"go-modules-api/internal/server/http/middleware"
	"go-modules-api/internal/server/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

type Server struct {
	App *fiber.App
	Log *zap.Logger
}

func NewServer(log *zap.Logger) *Server {
	app := fiber.New(fiber.Config{
		AppName:   "go-modules-api",
		BodyLimit: 1024 * 1024 * 1024, // 1GB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(middleware.RequestLogger(log))

	app.Static("/", "./docs")

	appContainer := container.NewAppContainer()

	routes.SetupRoutes(app, log, appContainer)

	return &Server{
		App: app,
		Log: log,
	}
}

func (s *Server) Start() {
	port := fmt.Sprintf(":%d", config.Env.AppPort)
	log := s.Log.Named("server")
	log.Sugar().Infof("server is running at %s", port)
	if err := s.App.Listen(port); err != nil {
		log.Fatal("error while starting server", zap.Error(err))
	}
}
