package services

import (
	"note_add/note_add_api/api/repositories"
	"note_add/note_add_api/models"
)

type NoteContentServices struct {
	repo repositories.NoteContentRepository
}

func SaveNoteContentService(repo repositories.NoteContentRepository) NoteContentServices {
	return NoteContentServices{
		repo: repo,
	}
}

func (n NoteContentServices) SaveNoteContents(note models.SaveNoteContents) error {
	return n.repo.SaveNoteContents(note)
}

func (n NoteContentServices) UpdateNoteContents(note models.SaveNoteContents) error {
	return n.repo.UpdateNoteContents(note)
}

func (n NoteContentServices) GetNoteContents(note models.NoteContents) (*models.NoteContents, error) {
	return n.repo.GetNoteContents(note)
}
