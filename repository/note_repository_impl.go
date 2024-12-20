package repository

import (
	"github.com/tuco5/fiber-demo/helper"
	"github.com/tuco5/fiber-demo/model"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{
		Db: db,
	}
}

// Delete implements NoteRepository.
func (n *NoteRepositoryImpl) Delete(id int) {
	result := n.Db.Delete(&model.Note{}, id)
	helper.ErrorPanic(result.Error)
}

// FindAll implements NoteRepository.
func (n *NoteRepositoryImpl) FindAll() []model.Note {
	var notes []model.Note
	result := n.Db.Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}

// FindById implements NoteRepository.
func (n *NoteRepositoryImpl) FindById(id int) (model.Note, error) {
	var note model.Note
	result := n.Db.First(&note, id)
	if result.Error != nil {
		return note, result.Error
	}
	return note, nil
}

// Save implements NoteRepository.
func (n *NoteRepositoryImpl) Save(note model.Note) {
	result := n.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

// Update implements NoteRepository.
func (n *NoteRepositoryImpl) Update(note model.Note) {
	result := n.Db.Save(&note)
	helper.ErrorPanic(result.Error)
}
