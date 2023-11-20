package controller

import (
	// "fmt"
	// "os"
	"strconv"
	// "time"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateShop(c *gin.Context) {

	var shop models.Shop
	c.ShouldBindJSON(&shop)

	// Unique Email Validation
	var countEmail int64 = 0
	config.DB.Model(models.Shop{}).Where("email = ?", shop.ShopEmail).Count(&countEmail)

	if countEmail != 0 {
		c.String(200, "Email is Not Unique")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shop.ShopPassword), 10)

	if err != nil {
		c.String(200, "Password Hashing Failed")
		return
	}

	shop.ShopPassword = string(hashedPassword)

	config.DB.Create(&shop)
	c.JSON(200, &shop)

}

func GetShops(c *gin.Context) {

	type RequestBody struct {
		PageNumber int  `json:"page_number"`
		IsActive   bool `json:"is_active"`
		IsBanned   bool `json:"is_banned"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	pageSize := 4

	if requestBody.IsActive && requestBody.IsBanned {

		shops := []models.Shop{}
		config.DB.Model(models.Shop{}).Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&shops)

		c.JSON(200, shops)
		return

	}

	if requestBody.IsActive {

		shops := []models.Shop{}
		config.DB.Model(models.Shop{}).Where("status = ?", "Active").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&shops)

		c.JSON(200, shops)
		return

	}

	if requestBody.IsBanned {

		shops := []models.Shop{}
		config.DB.Model(models.Shop{}).Where("status = ?", "Banned").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&shops)

		c.JSON(200, shops)
		return

	}

	shops := []models.Shop{}
	c.JSON(200, &shops)

}

func GetTopShops(c *gin.Context) {

	type RequestBody struct {
		Limit int `json:"limit"`
	}
	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Top With Most Items Sold
	// query := `
	// 	SELECT SUM(quantity), shop_id
	// 	FROM (
	// 		SELECT OD.product_id,
	// 			SUM(quantity) AS quantity,
	// 			shop_id
	// 		FROM order_details OD JOIN products PS ON
	// 			OD.product_id = PS.product_id
	// 		GROUP BY OD.product_id, shop_id
	// 	) AS sub
	// 	GROUP BY shop_id
	// 	ORDER BY SUM(quantity) DESC
	// 	LIMIT ` + strconv.Itoa(requestBody.Limit)


	query := `SELECT DISTINCT shop_id
	FROM products;`+strconv.Itoa(requestBody.Limit)
	rows, _ := config.DB.Raw(query).Rows()

	type Result struct {
		// Sum    int  `json:"sum"`
		ShopID uint `json:"shop_id"`
	}

	var shopIds []uint

	for rows.Next() {

		var row Result
		err := rows.Scan(&row.ShopID)
		if err != nil {
			panic(err)
		}

		shopIds = append(shopIds, row.ShopID)

	}

	var shops []models.Shop
	config.DB.Model(models.Shop{}).Where("id IN ?", shopIds).Find(&shops)

	c.JSON(200, shops)

}
