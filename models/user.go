package models

import (
	"time"

	"gorm.io/gorm"
)

// User Model
type Users struct {
	ID        int64          `gorm:"primaryKey;auto_incriment" json:"id"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Returns the table name of the Model
func (users *Users) TableName() string {
	return "users"
}

// Request Binding for User Login
type UserLogin struct {
	Email    string `form:"email"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

// Request Binding  for User Registration
type UserRegistration struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (user *Users) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["email"] = user.Email
	resp["password"] = user.Password
	resp["username"] = user.Username
	resp["is_active"] = user.IsActive

	return resp
}
