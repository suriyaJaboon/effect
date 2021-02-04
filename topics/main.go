package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	f := fiber.New()
	f.Use(
		cors.New(),
		logger.New(),
		requestid.New(),
	)

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello I'm fiber!")
	})

	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{"code": "OK", "message": "I'm fiber."})
	})
	log.Fatal(f.Listen(":3000"))
}
