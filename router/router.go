package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuco5/fiber-demo/controller"
)

func NewNoteRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to fiber demo",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:id", func(router fiber.Router) {
		router.Get("", noteController.FindById)
		router.Put("", noteController.Update)
		router.Delete("", noteController.Delete)
	})

	return router
}
