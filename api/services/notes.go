package services

import (
	"note_add/note_add_api/api/repositories"
	"note_add/note_add_api/models"
)

type NotesService struct {
	repo repositories.NotesRepository
}

func SaveNoteService(repo repositories.NotesRepository) NotesService {
	return NotesService{
		repo: repo,
	}
}

func (n NotesService) SaveNotes(note models.SaveNotes) error {
	return n.repo.SaveNotes(note)
}

func (n NotesService) UpdateNote(note models.SaveNotes) error {
	return n.repo.UpdateNotes(note)
}

func (n NotesService) GetNotes(note models.Notes, noteId string) []models.Notes {
	return n.repo.GetNotes(note, noteId)
}
