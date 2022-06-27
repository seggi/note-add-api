package routes

import (
	"note_add/note_add_api/api/controllers"
	"note_add/note_add_api/infrastructure"
)

type NotesRoute struct {
	Handler    infrastructure.GinRouter
	Controller controllers.NotesController
}

func SaveNoteRoute(
	controller controllers.NotesController,
	handler infrastructure.GinRouter) NotesRoute {
	return NotesRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (n NotesRoute) Setup() {
	note := n.Handler.Gin.Group("/note")
	{
		note.POST("/record-new-note", n.Controller.SaveNotes)
	}
}
