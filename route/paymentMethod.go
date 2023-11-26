package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func PaymentMethodRoute(router *gin.Engine) {

	router.GET("/get-payment-methods", controller.GetPaymentMethods)

}
