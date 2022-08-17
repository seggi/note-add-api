package utils

import "github.com/gin-gonic/gin"

// Response  struct
type Response struct {
	Success     bool              `json:"success"`
	Message     string            `json:"message"`
	Data        map[string]string `json:"data"`
	AccessToken string            `json:"access_token"`
}

// Json Error Response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

// Json  Success Response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
