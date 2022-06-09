package routes

import (
	"note_add/note_add_api/api/controllers"
	"note_add/note_add_api/infrastructure"
)

// Route for user module
type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controllers.UserController
}

// Initialized new instance of UserRoute
func NoteAddUserRoute(
	controller controllers.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

// Setup User routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.Login)
	}
}
