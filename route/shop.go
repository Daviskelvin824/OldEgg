package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func ShopRoute(router *gin.Engine) {

	router.POST("/create-shop", controller.CreateShop)
	router.POST("/get-top-shops", controller.GetTopShops)

}
