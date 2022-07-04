package routes

import (
	"note_add/note_add_api/api/controllers"
	"note_add/note_add_api/infrastructure"
)

type NoteContentRoute struct {
	Handler    infrastructure.GinRouter
	Controller controllers.NotesContentController
}

func SaveNoteContentRoute(
	controller controllers.NotesContentController,
	handler infrastructure.GinRouter) NoteContentRoute {
	return NoteContentRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (n NoteContentRoute) Setup() {
	note := n.Handler.Gin.Group("/note-content")
	{
		// note.Use(middlewares.JwtAuthMiddleware())
		note.POST("/record-note-content", n.Controller.SaveNoteContents)
		note.PUT("/update-note-content", n.Controller.UpdateNoteContents)
		note.GET("/get-note-content", n.Controller.GetNoteContents)
	}
}
