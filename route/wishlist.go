package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func WishlistRoute(router *gin.Engine) {

	router.POST("/create-wishlist", controller.CreateWishlist)
	router.POST("/get-wishlists", controller.GetWishlists)
	router.POST("/update-wishlist", controller.UpdateWishlist)
	router.POST("/update-item-quantity", controller.UpdateItemQuantity)

	router.POST("/get-wishlists-with-details", controller.GetWishlistsWithDetails)
	router.POST("/get-wishlist-details-by-id", controller.GetWishlistDetailsByID)

	router.POST("/add-to-wishlist", controller.AddToWishlist)
	router.POST("/remove-from-wishlist", controller.RemoveFromWishlist)
	router.POST("/save-wishlist-note", controller.SaveWishlistNote)

	router.POST("/add-wishlist-items-to-cart", controller.AddWishlistItemsToCart)
	router.POST("/add-all-items-to-cart", controller.AddAllItemsToCart)

	router.POST("/get-public-wishlists", controller.GetAllPublicWishlists)

	router.POST("/duplicate-wishlist", controller.DuplicateWishlist)

}
