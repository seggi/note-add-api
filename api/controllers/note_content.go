package controllers

import (
	"net/http"
	"note_add/note_add_api/api/services"
	"note_add/note_add_api/models"
	"note_add/note_add_api/utils"

	"github.com/gin-gonic/gin"
)

type NotesContentController struct {
	services services.NoteContentServices
}

func SaveNotesContentController(s services.NoteContentServices) NotesContentController {
	return NotesContentController{
		services: s,
	}
}

func (n *NotesContentController) SaveNoteContents(c *gin.Context) {
	var noteContent models.SaveNoteContents

	if err := c.ShouldBind(&noteContent); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	n.services.SaveNoteContents(noteContent)
	utils.SuccessJSON(c, http.StatusOK, "Note title saved with success")

}

func (n *NotesContentController) UpdateNoteContents(c *gin.Context) {
	var note models.SaveNoteContents

	if err := c.ShouldBind(&note); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	err := n.services.UpdateNoteContents(note)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error updating Note title")
		return
	}
	utils.SuccessJSON(c, http.StatusOK, "Note title updated with success")
}
