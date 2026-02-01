package services

import (
	"errors"
	"notesmanagement/internal/models"
	"notesmanagement/internal/repositories"
)

type NoteService struct{
	repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService{
	return &NoteService{repo: repo}
}

func (s *NoteService) GetAll() ([]models.Note, error){
	return s.repo.GetAll()
}

func (s *NoteService) Create(data *models.Note) error{
	if data.Title == "" {
		return errors.New("title wajib diisi")
	}

	if data.CategoryID == 0 {
		data.CategoryID = 1
	}
	
	if data.CategoryID < 1 {
		return errors.New("request tidak sesuai tabel categories")
	}

	return s.repo.Create(data)
}

func (s *NoteService) GetByID(id int) (*models.Note, error){
	return s.repo.GetByID(id)
}

func (s *NoteService) Update(note *models.Note) error{
	return  s.repo.Update(note)
}

func (s *NoteService) Delete(id int) error{
	return s.repo.Delete(id)
}
