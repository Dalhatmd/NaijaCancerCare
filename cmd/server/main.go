package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/dalhatmd/NaijaCancerCare/internal/config"
)

func main() {
	app := fiber.New()
	
	cfg := config.Load()


	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(cfg.AppPort))
}
