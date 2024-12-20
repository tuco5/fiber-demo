package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tuco5/fiber-demo/data/request"
	"github.com/tuco5/fiber-demo/data/response"
	"github.com/tuco5/fiber-demo/helper"
	"github.com/tuco5/fiber-demo/service"
)

type NoteController struct {
	service service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{
		service: service,
	}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	req := request.CreateNoteRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)

	controller.service.Create(req)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Success create note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	req := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)

	id := ctx.Params("id")
	req.Id, err = strconv.Atoi(id)
	helper.ErrorPanic(err)

	controller.service.Update(req)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Success update note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	controller.service.Delete(id)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Success delete note",
		Data:    nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("id")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	data := controller.service.FindById(id)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Success find note",
		Data:    data,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	data := controller.service.FindaAll()

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Success find all note",
		Data:    data,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
