package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func SaveItemForLater(c *gin.Context) {

	var item models.SavedForLaterItems
	c.ShouldBindJSON(&item)
	config.DB.Model(models.SavedForLaterItems{}).Create(&item)
	c.JSON(200, item)

}

func GetSavedForLaterItems(c *gin.Context) {

	type RequestBody struct {
		UserID int64 `json:"user_id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var items []models.SavedForLaterItems
	config.DB.Model(models.SavedForLaterItems{}).Where("user_id = ?", &requestBody.UserID).Find(&items)

	type Response struct {
		Item    models.SavedForLaterItems `json:"saved_for_later_item"`
		Product models.Product            `json:"product"`
	}

	var response []Response

	length := len(items)
	for i := 0; i < length; i++ {

		var entry Response
		entry.Item = items[i]
		config.DB.Model(models.Product{}).Where("product_id = ?", items[i].ProductID).Find(&entry.Product)

		response = append(response, entry)

	}

	c.JSON(200, response)
}

