package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {

	type RequestBody struct {
		UserID    int64  `json:"user_id"`
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var entry models.Cart
	entry.UserID = int(requestBody.UserID)
	entry.ProductID = requestBody.ProductID
	entry.Quantity = requestBody.Quantity
	
	var existingCart models.Cart
	config.DB.Model(models.Cart{}).Where("user_id = ?", requestBody.UserID).Where("product_id = ?", requestBody.ProductID).First(&existingCart)

	if existingCart.ID == 0 {

		config.DB.Model(models.Cart{}).Create(&entry)

	} else {

		existingCart.Quantity += requestBody.Quantity
		config.DB.Save(&existingCart)

	}

	c.String(200, "Item Added to Cart Successfully")

}

func GetItemsInCart(c *gin.Context) {

	type RequestBody struct {
		UserID int64 `json:"user_id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var carts []models.Cart
	config.DB.Model(models.Cart{}).Where("user_id = ?", requestBody.UserID).Find(&carts)

	type ResponseBody struct {
		UserID             int      `json:"user_id" gorm:"References:users(ID)"`
		ProductID          string   `json:"product_id"`
		Quantity           int      `json:"quantity"`
		ShopID             int      `json:"shop_id" gorm:"References:users(ID)"`
		ProductCategoryID  int      `json:"product_category_id" gorm:"References:product_categories(ProductCategoryID)"`
		ProductName        string   `json:"product_name"`
		ProductDescription string   `json:"product_description"`
		ProductPrice       float64  `json:"product_price"`
		ProductStock       int      `json:"product_stock"`
		ProductDetails     string   `json:"product_details"`
		ProductImageLinks  string `json:"product_link"`
	}

	var products []ResponseBody
	length := len(carts)
	for i := 0; i < length; i++ {

		var product models.Product
		config.DB.Model(models.Product{}).Where("product_id = ?", carts[i].ProductID).First(&product)

		var response ResponseBody
		response.UserID = carts[i].UserID
		response.ProductID = carts[i].ProductID
		response.Quantity = carts[i].Quantity
		response.ShopID = product.ShopID
		response.ProductCategoryID = product.ProductCategoryID
		response.ProductName = product.ProductName
		response.ProductDescription = product.ProductDescription
		response.ProductPrice = product.ProductPrice
		response.ProductStock = product.ProductStock
		response.ProductDetails = product.ProductDetails
		response.ProductImageLinks = product.ProductLink
		products = append(products, response)

	}

	c.JSON(200, products)

}

func UpdateItemsInCart(c *gin.Context) {

	type RequestBody struct {
		UserID    int    `json:"user_id" gorm:"References:users(ID)"`
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var cart models.Cart
	config.DB.Model(models.Cart{}).Where("user_id = ?", requestBody.UserID).Where("product_id = ?", requestBody.ProductID).First(&cart)

	if cart.ID == 0 {
		c.String(200, "Cart Item Not Found")
		return
	}

	cart.Quantity = requestBody.Quantity
	config.DB.Save(&cart)

	c.String(200, "Item Saved Successfully")

}

func RemoveFromCart(c *gin.Context) {

	var cart models.Cart
	c.ShouldBindJSON(&cart)

	var toDelete models.Cart
	config.DB.Model(models.Cart{}).Where("user_id = ?", cart.UserID).Where("product_id = ?", cart.ProductID).First(&toDelete)

	config.DB.Model(models.Cart{}).Delete(&toDelete)

	c.String(200, "Item Deleted From Cart")

}
