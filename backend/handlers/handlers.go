package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/generateJD-microservices.git/model"
	"github.com/huuloc2026/generateJD-microservices.git/services/latex"
)

var generator latex.Generator = latex.NewHTTPClient()

func GenerateHandler(c *fiber.Ctx) error {
	var req model.JDRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	pdfBytes, err := generator.Generate(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "PDF generation failed"})
	}

	c.Set("Content-Type", "application/pdf")
	return c.Send(pdfBytes)
}
