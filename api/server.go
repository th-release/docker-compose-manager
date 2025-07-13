package api

import (
	"th-release/dcm/api/docker"
	"th-release/dcm/utils"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App    *fiber.App
	Config utils.Config
}

func InitServer(config *utils.Config) *ServerConfig {
	app := fiber.New()

	if config == nil {
		return nil
	}

	server := &ServerConfig{
		App:    app,
		Config: *config,
	}

	server.setupRoutes()
	return server

}

func (s *ServerConfig) setupRoutes() {
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("<h1>Docker Compose Manager</h1>")
	})

	apiGroup := s.App.Group("/api", ApiMiddleware)

	apiGroup.Post("/insert", docker.Insert)
	apiGroup.Delete("/delete", docker.Delete)
}
