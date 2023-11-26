package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func GetDeliveryProviders(c *gin.Context) {

	var deliveryProviders []models.DeliveryProvider
	config.DB.Model(models.DeliveryProvider{}).Find(&deliveryProviders)
	c.JSON(200, deliveryProviders)

}
