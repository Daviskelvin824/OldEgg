package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func GetProductCategories(c *gin.Context) {

	productCategories := []models.ProductCategory{}
	config.DB.Find(&productCategories)
	c.JSON(200, &productCategories)

}

func CreateProductCategory(c *gin.Context) {
	var productCategory models.ProductCategory
	if err := c.ShouldBindJSON(&productCategory); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	config.DB.Create(&productCategory)

	c.JSON(201, productCategory)
}

func GetPopularProductCategories(c *gin.Context) {

	type Result struct {
		ProductCategoryName string `json:"product_category_name"`
	}
	var result []Result

	rows, _ := config.DB.Raw(`SELECT product_category_name
				FROM product_categories
					JOIN products ON 
						product_categories.product_category_id = products.product_category_id
				GROUP BY product_categories.product_category_id,
						product_category_name
				LIMIT 5`).Rows()

	for rows.Next() {

		var row Result
		err := rows.Scan(&row.ProductCategoryName)
		if err != nil {
			panic(err)
		}

		result = append(result, row)

	}

	c.JSON(200, result)

}


func GetProductCategoryByShopID(c *gin.Context) {

	type RequestBody struct {
		ShopID int64 `json:"id"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var productCategoryIDs []int64
	config.DB.Model(models.Product{}).Where("shop_id = ?", requestBody.ShopID).Distinct().Pluck("product_category_id", &productCategoryIDs)

	var productCategories []models.ProductCategory
	config.DB.Model(models.ProductCategory{}).Where("id IN ?", productCategoryIDs).Find(&productCategories)

	c.JSON(200, productCategories)

}
