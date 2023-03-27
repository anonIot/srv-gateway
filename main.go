package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SRV Mini Gateway API Service")
	})

	app.Listen(":3033")

}
