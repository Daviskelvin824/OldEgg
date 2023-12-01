package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func ProductCategoryRoute(router *gin.Engine) {

	router.GET("/get-product-categories", controller.GetProductCategories)
	router.POST("/get-product-category-by-shop-id", controller.GetProductCategoryByShopID)
	router.GET("/get-popular-product-categories", controller.GetPopularProductCategories)

}

