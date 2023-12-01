package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {

	var address models.Address
	c.ShouldBindJSON(&address)

	config.DB.Model(models.Address{}).Create(&address)
	c.JSON(200, address)

}

func GetAddresses(c *gin.Context) {

	type RequestBody struct {
		UserID int64 `json:"user_id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var addresses []models.Address
	config.DB.Model(models.Address{}).Where("user_id = ?", requestBody.UserID).Find(&addresses)

	c.JSON(200, addresses)

}

func RemoveAddress(c *gin.Context) {

	var address, toDelete models.Address
	c.ShouldBindJSON(&address)

	config.DB.Model(models.Address{}).Where("id = ?", address.ID).Find(&toDelete)
	config.DB.Model(models.Address{}).Delete(&toDelete)

	c.String(200, "Address Deleted")

}
