package main

import (
	"errors"
	"note_add/note_add_api/api/controllers"
	"note_add/note_add_api/api/repositories"
	"note_add/note_add_api/api/routes"
	"note_add/note_add_api/api/services"
	"note_add/note_add_api/infrastructure"
	"note_add/note_add_api/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NoteAddRouter()
	db := infrastructure.NoteAddDatabase()

	postRepository := repositories.NoteAddUserRepository(db)
	postService := services.NoteAddUserService(postRepository)
	postController := controllers.NoteAddUserController(postService)
	postRoute := routes.NoteAddUserRoute(postController, router)
	postRoute.Setup()

	// Check if table exists
	if err := db.DB.AutoMigrate(&models.User{}, &models.Notes{}, &models.NoteContents{}); err != nil {
		errors.New("Unable autoMigrateDB - " + err.Error())
	}

	router.Gin.Run(":7000")
}
