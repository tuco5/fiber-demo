package repository

import "github.com/tuco5/fiber-demo/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(id int)
	FindById(id int) (model.Note, error)
	FindAll() []model.Note
}
