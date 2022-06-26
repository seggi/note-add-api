package models

import (
	"time"
)

// Note Model

type Notes struct {
	ID        int64     `gorm:"primary_key;auto_incriment" json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

func (notes *Notes) TableName() string {
	return "notes"
}

type SaveNotes struct {
	UserID int    `form:"userID" binding:"required"`
	Title  string `form:"title" binding:"required"`
}

func (notes *Notes) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = notes.ID
	resp["userId"] = notes.UserID
	resp["title"] = notes.Title

	return resp
}
