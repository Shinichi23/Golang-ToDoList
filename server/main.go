package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, world!")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	log.Fatal(app.Listen(":3000"))
}
