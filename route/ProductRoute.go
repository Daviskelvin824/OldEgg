package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {

	// router.POST("/create-product", controller.CreateProduct)
	// router.POST("/get-products", controller.GetProducts)
	// router.POST("/get-product-by-id", controller.GetProductByID)
	// router.POST("/update-product", controller.UpdateProduct)
	// router.POST("/search-product", controller.SearchProduct)

	router.POST("/get-recommended-products", controller.GetRecommendedProducts)
	// router.POST("/get-recommended-products-by-shop", controller.GetRecommendedProductsByShop)

}
