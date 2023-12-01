package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.Engine) {

	router.POST("add-to-cart", controller.AddToCart)
	router.POST("get-items-in-cart", controller.GetItemsInCart)
	router.POST("update-items-in-cart", controller.UpdateItemsInCart)
	router.POST("remove-from-cart", controller.RemoveFromCart)

}
