package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	port := 3000
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("HELLO")
	})

	log.Printf("Server started on port %d", port)

	app.Listen(fmt.Sprintf(":%d", port))
}
