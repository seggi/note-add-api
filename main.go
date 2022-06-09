package main

import (
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

	db.DB.AutoMigrate(&models.User{})

	router.Gin.Run(":7000")
}
