package main

import (
	"github.com/gofiber/fiber/v2"
)

func handle(c *fiber.Ctx) error {
	return c.SendString("Welcome Dr. Vipin Classes")
}
func main() {
	// Prefork will spawn child processes
	app := fiber.New()
	app.Use("/vipin", handle)
	app.Static("/", "./template")
	//app.Get("/", handle)
	app.Listen(":9999")
}
