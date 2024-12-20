package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tuco5/fiber-demo/config"
	"github.com/tuco5/fiber-demo/controller"
	"github.com/tuco5/fiber-demo/model"
	"github.com/tuco5/fiber-demo/repository"
	"github.com/tuco5/fiber-demo/router"
	"github.com/tuco5/fiber-demo/service"
)

func main() {
	fmt.Println("Starting server...")

	// Load env variables
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to DB
	db := config.ConnectDB(&env)
	db.Table("notes").AutoMigrate(&model.Note{})

	// Init repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init service
	validate := validator.New()
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Init controller
	noteController := controller.NewNoteController(noteService)

	// Init router
	router_v1 := router.NewNoteRouter(noteController)

	// Init fiber app and mount router
	app := fiber.New()
	api := app.Group("/api")
	api.Mount("/v1", router_v1)

	// Start listening
	fmt.Println("Server started on port 3000")
	app.Listen((":3000"))
}
