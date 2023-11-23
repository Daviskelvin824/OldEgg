package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func GetOneTimeCode(c *gin.Context) {

	type RequestBody struct {
		Email string `json:"email"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	oneTimeCode := []models.OneTimeCode{}
	config.DB.Where("email = ?", requestBody.Email).First(&oneTimeCode)
	c.JSON(200, &oneTimeCode)

}
