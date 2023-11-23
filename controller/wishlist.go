package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateWishlist(c *gin.Context) {

	var wishlist models.Wishlist
	c.ShouldBindJSON(&wishlist)
	config.DB.Model(models.Wishlist{}).Create(&wishlist)
	c.JSON(200, wishlist)

}

func GetWishlists(c *gin.Context) {

	type RequestBody struct {
		UserID int64 `json:"user_id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var wishlists []models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("user_id = ?", requestBody.UserID).Find(&wishlists)

	c.JSON(200, wishlists)

}

// func GetWishlistsWithDetails(c *gin.Context) {

// 	type RequestBody struct {
// 		UserID int64 `json:"user_id"`
// 	}

// 	var requestBody RequestBody
// 	c.ShouldBindJSON(&requestBody)

// 	var wishlists []models.Wishlist
// 	config.DB.Model(models.Wishlist{}).Where("user_id = ?", requestBody.UserID).Find(&wishlists)

// 	type Response struct {
// 		Header             models.Wishlist         `json:"header"`
// 		Detail             []models.WishlistDetail `json:"detail"`
// 		Products           []models.Product        `json:"products"`
// 		NumberOfPromotions int                    `json:"number_of_promotions"`
// 	}

// 	var responses []Response

// 	length := len(wishlists)
// 	for i := 0; i < length; i++ {

// 		var response Response

// 		var details []models.WishlistDetail
// 		config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", wishlists[i].ID).Find(&details)

// 		response.Header = wishlists[i]
// 		response.Detail = details
// 		response.NumberOfPromotions = 0

// 		detailsLength := len(details)
// 		for j := 0; j < detailsLength; j++ {

// 			var promotion models.ProductPromotion
// 			config.DB.Model(models.ProductPromotion{}).Where("product_id = ?", details[j].ProductID).First(&promotion)
// 			if promotion.ID != 0 {
// 				response.NumberOfPromotions++
// 			}

// 			var product models.Product
// 			config.DB.Model(models.Product{}).Where("product_id = ?", details[j].ProductID)
// 			response.Products = append(response.Products, product)

// 		}

// 		responses = append(responses, response)

// 	}

// 	c.JSON(200, responses)

// }

func GetWishlistDetailsByID(c *gin.Context) {

	type RequestBody struct {
		WishlistID uint `json:"wishlist_id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var wishlist models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("id = ?", requestBody.WishlistID).First(&wishlist)

	type Response struct {
		Header  models.Wishlist       `json:"header"`
		Detail  models.WishlistDetail `json:"detail"`
		Product models.Product        `json:"product"`
	}

	var responses []Response

	var header models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("id = ?", wishlist.ID).First(&header)

	var details []models.WishlistDetail
	config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", wishlist.ID).Find(&details)

	detailsLength := len(details)
	for i := 0; i < detailsLength; i++ {

		var response Response

		var product models.Product
		config.DB.Model(models.Product{}).Where("product_id = ?", details[i].ProductID).First(&product)

		response.Header = header
		response.Detail = details[i]
		response.Product = product

		responses = append(responses, response)

	}

	if detailsLength == 0 {

		var response Response
		response.Header = header
		responses = append(responses, response)

	}

	c.JSON(200, responses)

}

func AddToWishlist(c *gin.Context) {

	var wishlistDetail models.WishlistDetail
	c.ShouldBindJSON(&wishlistDetail)

	config.DB.Model(models.WishlistDetail{}).Create(&wishlistDetail)
	c.JSON(200, wishlistDetail)

}

func RemoveFromWishlist(c *gin.Context) {

	var requestBody models.WishlistDetail
	c.ShouldBindJSON(&requestBody)

	var toDelete models.WishlistDetail
	config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", requestBody.WishlistDetailID).Where("product_id = ?", requestBody.ProductID).Find(&toDelete)

	config.DB.Model(models.WishlistDetail{}).Delete(&toDelete)

	c.JSON(200, toDelete)

}

func UpdateWishlist(c *gin.Context) {

	type RequestBody struct {
		Token        string `json:"token"`
		WishlistID   uint   `json:"wishlist_id"`
		WishlistName string `json:"wishlist_name"`
		IsPrivate    bool   `json:"is_private"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Only Update if Token is Valid
	tokenString := requestBody.Token
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRETKEY")), nil

	})

	if err != nil {
		c.String(200, "Token Parsing Failed")
		return
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {

		if float64(time.Now().Unix()) > claims["expire"].(float64) {

			c.String(200, "Cookie Expired")
			return

		}

		var user models.User
		config.DB.First(&user, "email = ?", claims["subject"])

		var wishlist models.Wishlist
		config.DB.Model(&models.Wishlist{}).Where("id = ?", requestBody.WishlistID).First(&wishlist)

		if user.ID == uint(wishlist.UserID) {

			var toUpdate models.Wishlist
			config.DB.Model(models.Wishlist{}).Where("id = ?", requestBody.WishlistID).First(&toUpdate)

			toUpdate.WishListName = requestBody.WishlistName
			toUpdate.IsPrivate = requestBody.IsPrivate

			config.DB.Save(&toUpdate)

			c.String(200, "Wishlist Updated")

		} else {

			c.String(200, "You Are Not Authorized to Update This Product")

		}

	}

}

func SaveWishlistNote(c *gin.Context) {

	type Body struct {
		WishlistID int    `json:"wishlist_id"`
		Note       string `json:"note"`
	}
	var body Body
	c.ShouldBindJSON(&body)

	var header models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("id = ?", body.WishlistID).First(&header)

	header.Note = body.Note

	config.DB.Save(&header)

	c.String(200, "Note Saved")

}

func UpdateItemQuantity(c *gin.Context) {

	type Body struct {
		WishlistDetailID int    `json:"wishlist_detail_id"`
		ProductID        string `json:"product_id"`
		Quantity         int    `json:"quantity"`
	}
	var body Body
	c.ShouldBindJSON(&body)

	var wishlistDetail models.WishlistDetail
	config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", body.WishlistDetailID).Where("product_id = ?", body.ProductID).First(&wishlistDetail)

	wishlistDetail.Quantity = body.Quantity

	config.DB.Save(&wishlistDetail)

	c.String(200, "Quantity Updated")

}

func AddWishlistItemsToCart(c *gin.Context) {

	type Body struct {
		WishlistDetailID int `json:"wishlist_detail_id"`
	}
	var body Body
	c.ShouldBindJSON(&body)

}

func AddAllItemsToCart(c *gin.Context) {

	type Body struct {
		UserID           uint `json:"user_id"`
		WishlistDetailID uint `json:"wishlist_detail_id"`
	}
	var body Body
	c.ShouldBindJSON(&body)

	var wishlistDetails []models.WishlistDetail
	config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", body.WishlistDetailID).Find(&wishlistDetails)

	length := len(wishlistDetails)
	for i := 0; i < length; i++ {

		var entry models.Cart
		entry.UserID = int(body.UserID)
		entry.ProductID = wishlistDetails[i].ProductID
		entry.Quantity = int(wishlistDetails[i].Quantity)

		// If Exist Add, if not create
		var existingCart models.Cart
		config.DB.Model(models.Cart{}).Where("user_id = ?", body.UserID).Where("product_id = ?", entry.ProductID).First(&existingCart)

		if existingCart.ID == 0 {

			config.DB.Model(models.Cart{}).Create(&entry)

		} else {

			existingCart.Quantity += entry.Quantity
			config.DB.Save(&existingCart)

		}

		c.String(200, "Item Added to Cart Successfully")

	}

}

func GetAllPublicWishlists(c *gin.Context) {

	type Body struct {
		PageNumber int `json:"page_number"`
		PageSize   int `json:"page_size"`
	}
	var body Body
	c.ShouldBindJSON(&body)

	var wishlists []models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("is_private = ?", false).Limit(body.PageSize).Offset((body.PageNumber - 1) * body.PageSize).Find(&wishlists)
	c.JSON(200, wishlists)

}

func DuplicateWishlist(c *gin.Context) {

	type Body struct {
		UserID     uint `json:"user_id"`
		WishlistID uint `json:"wishlist_id"`
	}
	var body Body
	c.ShouldBindJSON(&body)

	// Get Entry in wishlists
	var wishlistHeader models.Wishlist
	config.DB.Model(models.Wishlist{}).Where("id = ?", body.WishlistID).First(&wishlistHeader)

	if wishlistHeader.ID == 0 {
		c.String(200, "Wishlist Not Found")
		return
	}

	// Get Entries in wishlist_details
	var wishlistDetails []models.WishlistDetail
	config.DB.Model(models.WishlistDetail{}).Where("wishlist_detail_id = ?", body.WishlistID).Find(&wishlistDetails)

	// Duplicate
	var newWishlist models.Wishlist
	newWishlist.WishListName = wishlistHeader.WishListName
	newWishlist.IsPrivate = wishlistHeader.IsPrivate
	newWishlist.UserID = int(body.UserID)
	config.DB.Model(models.Wishlist{}).Create(&newWishlist)

	length := len(wishlistDetails)
	for i := 0; i < length; i++ {

		var newWishlistDetail models.WishlistDetail
		newWishlistDetail.WishlistDetailID = int64(newWishlist.ID)
		newWishlistDetail.ProductID = wishlistDetails[i].ProductID
		newWishlistDetail.Quantity = wishlistDetails[i].Quantity

		config.DB.Model(models.WishlistDetail{}).Create(&newWishlistDetail)

	}

	c.String(200, "Wishlist Duplicated")

}
