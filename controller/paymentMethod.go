package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func GetPaymentMethods(c *gin.Context) {

	var paymentMethods []models.PaymentMethod
	config.DB.Model(models.PaymentMethod{}).Find(&paymentMethods)
	c.JSON(200, paymentMethods)

}
