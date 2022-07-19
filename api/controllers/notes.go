package controllers

import (
	"net/http"
	"note_add/note_add_api/api/services"
	"note_add/note_add_api/models"
	"note_add/note_add_api/utils"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	service services.NotesService
}

func SaveNotesController(s services.NotesService) NotesController {
	return NotesController{
		service: s,
	}
}

func (n *NotesController) SaveNotes(c *gin.Context) {
	var note models.SaveNotes
	if err := c.ShouldBind(&note); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	err := n.service.SaveNotes(note)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Note title  already exist")
		return
	}

	utils.SuccessJSON(c, http.StatusOK, "Note title saved with success")
}

func (n *NotesController) UpdateNote(c *gin.Context) {
	var note models.SaveNotes

	if err := c.ShouldBind(&note); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	err := n.service.UpdateNote(note)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Error updating Note title")
		return
	}
	utils.SuccessJSON(c, http.StatusOK, "Note title updated with success")
}

func (n *NotesController) GetNotes(c *gin.Context) {
	var notes models.Notes

	if err := c.ShouldBindJSON(&notes); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	dbNotes, err := n.service.GetNotes(notes)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &dbNotes)
}
