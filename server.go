package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tuco5/fiber-demo/config"
)

func main() {
	fmt.Println("Starting server...")

	env, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB(&env)

	app := fiber.New()
	api := app.Group("/api")

	router_v1 := fiber.New()
	api.Mount("/v1", router_v1)

	router_v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen((":3000"))
}
