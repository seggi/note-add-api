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
}

func (noteContents *NoteContents) TableName() string {
	return "noteContents"
}

type SaveNoteContents struct {
	UserID      int    `form:"user_id" json:"user_id" binding:"required"`
	NoteID      int    `form:"note_id" json:"note_id" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	TextBody    string `form:"text_body" json:"text_body" binding:"required"`
}

func (noteContents *NoteContents) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = noteContents.ID
	resp["note_id"] = noteContents.NoteID
	resp["description"] = noteContents.Description
	resp["text_body"] = noteContents.TextBody

	return resp
}
