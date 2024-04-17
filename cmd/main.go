package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	port := 3000
	app := fiber.New()

	apiV1 := app.Group("/api/v1/todos")

	apiV1.Get("/", func(c fiber.Ctx) error {
		todos := []string{"HELLO", "WORLD", "TODOS", "API"}
		return c.JSON(fiber.Map{
			"data": todos,
		})
	})

	apiV1.Get("/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"data": id,
		})
	})

	log.Printf("Server started on port %d", port)

	app.Listen(fmt.Sprintf(":%d", port))
}
