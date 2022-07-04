package repositories

import (
	"note_add/note_add_api/infrastructure"
	"note_add/note_add_api/models"
)

type NoteContentRepository struct {
	db infrastructure.Database
}

type NoteContentResult struct {
	Description string
	TextBody    string
}

func SaveNoteContentsRepository(db infrastructure.Database) NoteContentRepository {
	return NoteContentRepository{
		db: db,
	}
}

func (n NoteContentRepository) SaveNoteContents(note models.SaveNoteContents) error {
	var dbNoteContent models.NoteContents
	dbNoteContent.NoteID = note.NoteID
	dbNoteContent.Description = note.Description
	dbNoteContent.TextBody = note.TextBody

	return n.db.DB.Create(&dbNoteContent).Error
}

func (n NoteContentRepository) UpdateNoteContents(note models.SaveNoteContents) error {
	var dbNoteContent models.NoteContents
	dbNoteContent.NoteID = note.NoteID
	dbNoteContent.Description = note.Description
	dbNoteContent.TextBody = note.TextBody

	n.db.DB.Model(&dbNoteContent).Where("note_id = ?", note.NoteID).Updates(map[string]interface{}{"description": note.Description, "text_body": note.TextBody})
	return nil
}

func (n NoteContentRepository) GetNoteContents(noteContent models.NoteContents) (*models.NoteContents, error) {
	var dbNoteContent models.NoteContents
	noteId := noteContent.NoteID

	err := n.db.DB.Joins("JOIN notes ON notes.id = noteContents.note_id").Where("notes.user_id = ?", noteId).Find(&dbNoteContent).Error

	if err != nil {
		return nil, err
	}

	return &dbNoteContent, err
}
