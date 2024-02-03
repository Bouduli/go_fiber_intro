package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.SendString("producs")
	})

	app.Get("/cars", func(c *fiber.Ctx) error {
		return c.SendString("cars")
	})

	app.Listen(":3000")
}
