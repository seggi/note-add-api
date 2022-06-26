package models

import "time"

// Note  Content Model
type NoteContents struct {
	ID          int64     `gorm:"primary_key;auto_incriment" json:"id"`
	NoteID      int       `json:"note_id"`
	Note        Notes     `gorm:"foreignKey:NoteID"`
	Description string    `json:"description"`
	TextBody    string    `json:"text_body"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

func (noteContents *NoteContents) TableName() string {
	return "noteContents"
}

type SaveNoteContents struct {
	NoteID      int    `form:"noteID" binding:"required"`
	Description string `form:"description"`
	TextBody    string `form:"text_body"`
}

func (noteContents *NoteContents) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = noteContents.ID
	resp["note_id"] = noteContents.NoteID
	resp["description"] = noteContents.Description
	resp["text_body"] = noteContents.TextBody

	return resp
}
