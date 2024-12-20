package main

import (
	"fmt"

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

	// Load config
	validate := validator.New()
	cfg := config.Load(validate)

	// Connect to DB
	db := config.ConnectDB(cfg)
	db.Table("notes").AutoMigrate(&model.Note{})

	// Init repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init service
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
	fmt.Printf("Server started on port %s\n", cfg.Port)
	app.Listen(cfg.Port)
}
