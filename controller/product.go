package controller

import (
	// "fmt"
	// "os"
	// "time"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
)

func CreateProduct(c *gin.Context) {

	var product models.Product
	c.ShouldBindJSON(&product)
	config.DB.Create(&product)
	c.JSON(200, &product)

}

// func GetProducts(c *gin.Context) {

// 	type RequestBody struct {
// 		ShopID          int    `json:"shop_id"`
// 		PageNumber      int    `json:"page_number"`
// 		IsAvailableOnly bool   `json:"is_available_only"`
// 		Keyword         string `json:"keyword"`
// 		InnerKeyword    string `json:"inner_keyword"`
// 	}

// 	var requestBody RequestBody
// 	c.ShouldBindJSON(&requestBody)

// 	pageSize := 50

// 	var count int64
// 	rawProducts := []models.Product{}
// 	if requestBody.IsAvailableOnly {

// 		config.DB.Model(models.Product{}).Where("shop_id = ?", requestBody.ShopID).Where("product_stock > 0").Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&rawProducts)

// 		config.DB.Model(models.Product{}).Where("shop_id = ?", requestBody.ShopID).Where("product_stock > 0").Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Count(&count)

// 	} else {

// 		config.DB.Model(models.Product{}).Where("shop_id = ?", requestBody.ShopID).Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&rawProducts)

// 		config.DB.Model(models.Product{}).Where("shop_id = ?", requestBody.ShopID).Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Count(&count)

// 	}

// 	type Product struct {
// 		ProductID          string   `json:"product_id"`
// 		ShopID             int      `json:"shop_id"`
// 		ProductCategoryID  int      `json:"product_category_id"`
// 		ProductName        string   `json:"product_name"`
// 		ProductDescription string   `json:"product_description"`
// 		ProductPrice       float64  `json:"product_price"`
// 		ProductStock       int      `json:"product_stock"`
// 		ProductDetails     string   `json:"product_details"`
// 		ProductImageLinks  []string `json:"product_image_links"`
// 	}

// 	length := len(rawProducts)
// 	var parsedProducts []Product

// 	for i := 0; i < length; i++ {

// 		var product Product
// 		product.ProductID = rawProducts[i].ProductID
// 		product.ShopID = rawProducts[i].ShopID
// 		product.ProductCategoryID = rawProducts[i].ProductCategoryID
// 		product.ProductName = rawProducts[i].ProductName
// 		product.ProductDescription = rawProducts[i].ProductDescription
// 		product.ProductPrice = rawProducts[i].ProductPrice
// 		product.ProductStock = rawProducts[i].ProductStock
// 		product.ProductDetails = rawProducts[i].ProductDetails

// 		// Get ALl Image Links
// 		var productImageLinks []models.ProductImageLink
// 		config.DB.Model(models.ProductImageLink{}).Where("product_id = ?", product.ProductID).Find(&productImageLinks)

// 		var imageLinks []string
// 		productImageLinksLength := len(productImageLinks)
// 		for j := 0; j < productImageLinksLength; j++ {

// 			imageLinks = append(imageLinks, productImageLinks[j].Link)

// 		}
// 		product.ProductImageLinks = imageLinks

// 		parsedProducts = append(parsedProducts, product)

// 	}

// 	type ResponseBody struct {
// 		ParsedProducts []Product `json:"products"`
// 		Count          int64     `json:"count"`
// 	}

// 	var response ResponseBody
// 	response.ParsedProducts = parsedProducts
// 	response.Count = count

// 	c.JSON(200, &response)

// }

// func GetProductByID(c *gin.Context) {

// 	type RequestBody struct {
// 		ProductID string `json:"product_id"`
// 	}

// 	var requestBody RequestBody
// 	c.ShouldBindJSON(&requestBody)

// 	var rawProduct models.Product
// 	config.DB.First(&rawProduct, "product_id = ?", requestBody.ProductID)

// 	if rawProduct.ID == 0 {
// 		c.String(200, "Product Not Found")
// 		return
// 	}

// 	// Validate Banned Shop
// 	var shop models.Shop
// 	config.DB.Model(models.Shop{}).Where("id = ?", rawProduct.ShopID).First(&shop)

// 	if shop.Status == "Banned" {
// 		c.String(200, "Shop is Banned")
// 		return
// 	}

// 	type Product struct {
// 		ProductID          string   `json:"product_id"`
// 		ShopID             int      `json:"shop_id"`
// 		ProductCategoryID  int      `json:"product_category_id"`
// 		ProductName        string   `json:"product_name"`
// 		ProductDescription string   `json:"product_description"`
// 		ProductPrice       float64  `json:"product_price"`
// 		ProductStock       int      `json:"product_stock"`
// 		ProductDetails     string   `json:"product_details"`
// 		ProductImageLinks  []string `json:"product_image_links"`
// 	}

// 	var product Product
// 	product.ProductID = rawProduct.ProductID
// 	product.ShopID = rawProduct.ShopID
// 	product.ProductCategoryID = rawProduct.ProductCategoryID
// 	product.ProductName = rawProduct.ProductName
// 	product.ProductDescription = rawProduct.ProductDescription
// 	product.ProductPrice = rawProduct.ProductPrice
// 	product.ProductStock = rawProduct.ProductStock
// 	product.ProductDetails = rawProduct.ProductDetails

// 	// Get ALl Image Links
// 	var productImageLinks []models.ProductImageLink
// 	config.DB.Model(models.ProductImageLink{}).Where("product_id = ?", product.ProductID).Find(&productImageLinks)

// 	var imageLinks []string
// 	productImageLinksLength := len(productImageLinks)
// 	for j := 0; j < productImageLinksLength; j++ {

// 		imageLinks = append(imageLinks, productImageLinks[j].Link)

// 	}
// 	product.ProductImageLinks = imageLinks

// 	c.JSON(200, product)

// }

// func UpdateProduct(c *gin.Context) {

// 	type RequestBody struct {
// 		Token   string        `json:"token"`
// 		Product models.Product `json:"product"`
// 	}

// 	var requestBody RequestBody
// 	c.ShouldBindJSON(&requestBody)

// 	tokenString := requestBody.Token
// 	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(os.Getenv("SECRET")), nil

// 	})

// 	if err != nil {
// 		c.String(200, "Token Parsing Failed")
// 		return
// 	}

// 	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {

// 		if float64(time.Now().Unix()) > claims["expire"].(float64) {

// 			c.String(200, "Cookie Expired")
// 			return

// 		}

// 		var shop models.Shop
// 		config.DB.First(&shop, "shop_email = ?", claims["subject"])

// 		var product model.Product
// 		config.DB.Model(&models.Product{}).Where("product_id = ?", requestBody.Product.ProductID).First(&product)

// 		if shop.ID == uint(product.ShopID) {

// 			config.DB.Model(&model.Product{}).Where("product_id = ?", requestBody.Product.ProductID).Updates(map[string]interface{}{
// 				"product_name":        requestBody.Product.ProductName,
// 				"product_description": requestBody.Product.ProductDescription,
// 				"product_price":       requestBody.Product.ProductPrice,
// 				"product_stock":       requestBody.Product.ProductStock,
// 				"product_details":     requestBody.Product.ProductDetails,
// 			})

// 			c.JSON(200, &requestBody)

// 		} else {

// 			c.String(200, "You Are Not Authorized to Update This Product")

// 		}

// 	}

// }

// func SearchProduct(c *gin.Context) {

// 	type RequestBody struct {
// 		Keyword         string `json:"keyword"`
// 		InnerKeyword    string `json:"inner_keyword"`
// 		PageNumber      int    `json:"page_number"`
// 		IsAvailableOnly bool   `json:"is_available_only"`
// 	}

// 	var requestBody RequestBody
// 	c.ShouldBindJSON(&requestBody)

// 	pageSize := 50

// 	rawProducts := []model.Product{}

// 	if requestBody.IsAvailableOnly {

// 		config.DB.Model(model.Product{}).Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Where("product_stock > 0").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&rawProducts)

// 	} else {

// 		config.DB.Model(model.Product{}).Where("product_name ILIKE ?", "%"+requestBody.Keyword+"%").Where("product_name ILIKE ?", "%"+requestBody.InnerKeyword+"%").Limit(pageSize).Offset((requestBody.PageNumber - 1) * pageSize).Find(&rawProducts)

// 	}

// 	type Product struct {
// 		ProductID          string   `json:"product_id"`
// 		ShopID             int      `json:"shop_id"`
// 		ProductCategoryID  int      `json:"product_category_id"`
// 		ProductName        string   `json:"product_name"`
// 		ProductDescription string   `json:"product_description"`
// 		ProductPrice       float64  `json:"product_price"`
// 		ProductStock       int      `json:"product_stock"`
// 		ProductDetails     string   `json:"product_details"`
// 		ProductImageLinks  []string `json:"product_image_links"`
// 	}

// 	length := len(rawProducts)
// 	var parsedProducts []Product

// 	for i := 0; i < length; i++ {

// 		var product Product
// 		product.ProductID = rawProducts[i].ProductID
// 		product.ShopID = rawProducts[i].ShopID
// 		product.ProductCategoryID = rawProducts[i].ProductCategoryID
// 		product.ProductName = rawProducts[i].ProductName
// 		product.ProductDescription = rawProducts[i].ProductDescription
// 		product.ProductPrice = rawProducts[i].ProductPrice
// 		product.ProductStock = rawProducts[i].ProductStock
// 		product.ProductDetails = rawProducts[i].ProductDetails

// 		// Get ALl Image Links
// 		var productImageLinks []model.ProductImageLink
// 		config.DB.Model(model.ProductImageLink{}).Where("product_id = ?", product.ProductID).Find(&productImageLinks)

// 		var imageLinks []string
// 		productImageLinksLength := len(productImageLinks)
// 		for j := 0; j < productImageLinksLength; j++ {

// 			imageLinks = append(imageLinks, productImageLinks[j].Link)

// 		}
// 		product.ProductImageLinks = imageLinks

// 		parsedProducts = append(parsedProducts, product)

// 	}

// 	c.JSON(200, &parsedProducts)

// }

func GetRecommendedProducts(c *gin.Context) {

	type RequestBody struct {
		BatchNumber int `json:"batch_number"`
	}
	var requestBody RequestBody

	c.ShouldBindJSON(&requestBody)

	if requestBody.BatchNumber >= 4 {
		return
	}

	batchSize := 8

	var rawProducts []models.Product
	config.DB.Model(models.Product{}).Order("product_stock DESC").Limit(batchSize).Offset(batchSize * (requestBody.BatchNumber - 1)).Find(&rawProducts)

	type Product struct {
		ProductID          string   `json:"product_id"`
		ShopID             int      `json:"shop_id"`
		ProductCategoryID  int      `json:"product_category_id"`
		ProductName        string   `json:"product_name"`
		ProductDescription string   `json:"product_description"`
		ProductPrice       float64  `json:"product_price"`
		ProductStock       int      `json:"product_stock"`
		ProductDetails     string   `json:"product_details"`
		ProductLinks  		string `json:"product_link"`
	}
	var parsedProducts []Product

	length := len(rawProducts)
	for i := 0; i < length; i++ {

		var product Product
		product.ProductID = rawProducts[i].ProductID
		product.ShopID = rawProducts[i].ShopID
		product.ProductCategoryID = rawProducts[i].ProductCategoryID
		product.ProductName = rawProducts[i].ProductName
		product.ProductDescription = rawProducts[i].ProductDescription
		product.ProductPrice = rawProducts[i].ProductPrice
		product.ProductStock = rawProducts[i].ProductStock
		product.ProductDetails = rawProducts[i].ProductDetails
		product.ProductLinks = rawProducts[i].ProductLink
		// var productImageLinks []models.ProductImageLink
		// config.DB.Model(models.ProductImageLink{}).Where("product_id = ?", product.ProductID).Find(&productImageLinks)

		// var imageLinks []string
		// productImageLinksLength := len(productImageLinks)
		// for j := 0; j < productImageLinksLength; j++ {

		// 	imageLinks = append(imageLinks, productImageLinks[j].Link)

		// }
		// product.ProductLinks = imageLinks

		parsedProducts = append(parsedProducts, product)

	}

	c.JSON(200, parsedProducts)

}

// func GetRecommendedProductsByShop(c *gin.Context) {

// 	type RequestBody struct {
// 		ShopID int `json:"shop_id"`
// 	}
// 	var requestBody RequestBody

// 	c.ShouldBindJSON(&requestBody)

// 	batchSize := 8

// 	var rawProducts []model.Product
// 	config.DB.Model(model.Product{}).Where("shop_id = ?", requestBody.ShopID).Order("product_stock DESC").Limit(batchSize).Find(&rawProducts)

// 	type Product struct {
// 		ProductID          string   `json:"product_id"`
// 		ShopID             int      `json:"shop_id"`
// 		ProductCategoryID  int      `json:"product_category_id"`
// 		ProductName        string   `json:"product_name"`
// 		ProductDescription string   `json:"product_description"`
// 		ProductPrice       float64  `json:"product_price"`
// 		ProductStock       int      `json:"product_stock"`
// 		ProductDetails     string   `json:"product_details"`
// 		ProductImageLinks  []string `json:"product_image_links"`
// 	}
// 	var parsedProducts []Product

// 	length := len(rawProducts)
// 	for i := 0; i < length; i++ {

// 		var product Product
// 		product.ProductID = rawProducts[i].ProductID
// 		product.ShopID = rawProducts[i].ShopID
// 		product.ProductCategoryID = rawProducts[i].ProductCategoryID
// 		product.ProductName = rawProducts[i].ProductName
// 		product.ProductDescription = rawProducts[i].ProductDescription
// 		product.ProductPrice = rawProducts[i].ProductPrice
// 		product.ProductStock = rawProducts[i].ProductStock
// 		product.ProductDetails = rawProducts[i].ProductDetails

// 		var productImageLinks []model.ProductImageLink
// 		config.DB.Model(model.ProductImageLink{}).Where("product_id = ?", product.ProductID).Find(&productImageLinks)

// 		var imageLinks []string
// 		productImageLinksLength := len(productImageLinks)
// 		for j := 0; j < productImageLinksLength; j++ {

// 			imageLinks = append(imageLinks, productImageLinks[j].Link)

// 		}
// 		product.ProductImageLinks = imageLinks

// 		parsedProducts = append(parsedProducts, product)

// 	}

// 	c.JSON(200, parsedProducts)

// }
