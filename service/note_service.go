package service

import (
	"github.com/tuco5/fiber-demo/data/request"
	"github.com/tuco5/fiber-demo/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(id int)
	FindById(id int) response.NoteResponse
	FindaAll() []response.NoteResponse
}
