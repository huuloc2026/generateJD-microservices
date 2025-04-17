package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/generateJD-microservices.git/config"
	"github.com/huuloc2026/generateJD-microservices.git/handlers"
)

func main() {
	config.LoadConfig()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Jake Onyx 👋 !")
	})
	app.Post("/generate", handlers.GenerateHandler)

	app.Listen(":" + config.AppConfig.ServerPort)
}
