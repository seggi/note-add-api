package repositories

import (
	"note_add/note_add_api/infrastructure"
	"note_add/note_add_api/models"
	"note_add/note_add_api/utils"
)

// Accessing database
type UserRepository struct {
	db infrastructure.Database
}

// Instance on UserRepository
func NoteAddUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// Save user to database
func (u UserRepository) CreateUser(user models.UserRegistration) error {
	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.Password = user.Password
	dbUser.Username = user.Username
	dbUser.IsActive = true

	err := u.db.DB.Where("email = ?", user.Email).Find(&dbUser).Error
	if err != nil {
		return err
	}
	return u.db.DB.Create(&dbUser).Error
}

// LoginUser  -> return user

func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {
	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := utils.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
