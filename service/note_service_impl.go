package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/tuco5/fiber-demo/data/request"
	"github.com/tuco5/fiber-demo/data/response"
	"github.com/tuco5/fiber-demo/helper"
	"github.com/tuco5/fiber-demo/model"
	"github.com/tuco5/fiber-demo/repository"
)

type NoteServiceImpl struct {
	Repository repository.NoteRepository
	validate   *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		Repository: noteRepository,
		validate:   validate,
	}
}

// Create implements NoteService.
func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)

	noteModel := model.Note{
		Content: note.Content,
	}
	n.Repository.Save(noteModel)
}

// Delete implements NoteService.
func (n *NoteServiceImpl) Delete(id int) {
	n.Repository.Delete(id)
}

// FindById implements NoteService.
func (n *NoteServiceImpl) FindById(id int) response.NoteResponse {
	data, err := n.Repository.FindById(id)
	helper.ErrorPanic(err)

	return response.NoteResponse{
		Id:      data.Id,
		Content: data.Content,
	}
}

// FindaAll implements NoteService.
func (n *NoteServiceImpl) FindaAll() []response.NoteResponse {
	result := n.Repository.FindAll()
	var notes []response.NoteResponse

	for _, note := range result {
		notes = append(notes, response.NoteResponse{
			Id:      note.Id,
			Content: note.Content,
		})
	}

	return notes
}

// Update implements NoteService.
func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	data, err := n.Repository.FindById(note.Id)
	helper.ErrorPanic(err)

	data.Content = note.Content
	n.Repository.Update(data)
}
