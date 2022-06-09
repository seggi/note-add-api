package services

import (
	"note_add/note_add_api/api/repositories"
	"note_add/note_add_api/models"
)

// UserService struct
type UserService struct {
	repo repositories.UserRepository
}

// New user service -> get injected user repo
func NoteAddUserService(repo repositories.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

// Save users entity
func (u UserService) CreateUser(user models.UserRegistration) error {
	return u.repo.CreateUser(user)
}

// Login  -> Gets validated user
func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
