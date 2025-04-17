package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/generateJD-microservices.git/model"
	"github.com/huuloc2026/generateJD-microservices.git/services"
)

func GenerateHandler(c *fiber.Ctx) error {
	var req model.JDRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	pdfBytes, err := services.GeneratePDF(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to generate PDF"})
	}

	c.Set("Content-Type", "application/pdf")
	return c.Send(pdfBytes)
}
