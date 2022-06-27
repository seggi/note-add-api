package controllers

import (
	"net/http"
	"note_add/note_add_api/api/services"
	"note_add/note_add_api/models"
	"note_add/note_add_api/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// User controller struct
type UserController struct {
	service services.UserService
}

// New user controller
func NoteAddUserController(s services.UserService) UserController {
	return UserController{
		service: s,
	}
}

// Create user -> calls CreateUser services for validated user
func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegistration
	if err := c.ShouldBind(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
	}

	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.SuccessJSON(c, http.StatusOK, "Successfully Created User")
}

// Login user: Generates JWT Token for validated user

func (u *UserController) Login(c *gin.Context) {

	var user models.UserLogin
	var hmacSampleSecret []byte

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}

	response := &utils.Response{
		Success: true,
		Message: "Token generated successfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}
