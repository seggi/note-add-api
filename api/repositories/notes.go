package repositories

import (
	"note_add/note_add_api/infrastructure"
	"note_add/note_add_api/models"
)

type NotesRepository struct {
	db infrastructure.Database
}

func SaveNotesRepository(db infrastructure.Database) NotesRepository {
	return NotesRepository{
		db: db,
	}
}

// Record new note title

func (n NotesRepository) SaveNotes(note models.SaveNotes) error {
	var dbNotes models.Notes
	dbNotes.UserID = note.UserID
	dbNotes.Title = note.Title

	err := n.db.DB.Where(map[string]interface{}{"title": note.Title, "user_id": note.UserID}).Find(&dbNotes).Error
	if err != nil {
		return err
	}

	return n.db.DB.Create(&dbNotes).Error
}

// Update note title name

func (n NotesRepository) UpdateNotes(note models.SaveNotes) error {
	var dbNotes models.Notes
	dbNotes.UserID = note.UserID
	dbNotes.Title = note.Title

	err := n.db.DB.Model(&dbNotes).Where("user_id = ?", note.UserID).Error
	if err != nil {
		return err
	}

	n.db.DB.Model(&dbNotes).Where("user_id = ?", note.UserID).Update("title", note.Title)
	return nil
}