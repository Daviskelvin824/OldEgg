package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {

	router.POST("/create-order", controller.CreateOrder)
	router.POST("/get-orders", controller.GetOrders)
	router.POST("/mark-order-as-finished", controller.MarkOrderAsFinished)

	router.POST("/get-user-orders", controller.GetUserOrders)

}
