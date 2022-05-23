package models

import "time"

// Note  Content Model
type NoteContent struct {
	ID          int64     `gorm:"primary_key;auto_incriment" json:"id"`
	NoteID      int       `json:"note_id"`
	Description string    `json:"description"`
	Text_body   string    `json:"text_body"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
