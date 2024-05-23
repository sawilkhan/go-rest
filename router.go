package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sawilkhan/go-rest/handlers"
)

func generateApp() *fiber.App {
	app := fiber.New()

	//ping to app for health check
	app.Get("/health", func(c *fiber.Ctx) error{
		return c.SendString("OK")
	})

	//create library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.TestHandler)

	return app
}