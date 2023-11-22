package controller

import (
	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func CreateSearchQuery(c *gin.Context) {

	var searchQuery models.SearchQuery
	c.ShouldBindJSON(&searchQuery)

	var countSearchQueries int64
	config.DB.Model(models.SearchQuery{}).Where("email = ?", searchQuery.Email).Count(&countSearchQueries)

	if countSearchQueries >= 10 {
		c.String(200, "You Already Have 10 Saved Search Queries")
		return
	}

	config.DB.Create(&searchQuery)
	c.JSON(200, &searchQuery)
	c.String(200, "Search Query Saved")

}

func GetPopularSearchQueries(c *gin.Context) {

	query := `
		SELECT keyword, 
			COUNT(keyword)
		FROM search_queries
		GROUP BY keyword
		ORDER BY COUNT(keyword) DESC
		LIMIT 5
	`

	rows, _ := config.DB.Raw(query).Rows()

	type Result struct {
		Keyword string `json:"keyword"`
		Count   int64  `json:"count"`
	}

	var result []Result

	for rows.Next() {

		var row Result
		err := rows.Scan(&row.Keyword, &row.Count)
		if err != nil {
			panic(err)
		}

		result = append(result, row)

	}

	c.JSON(200, result)

}
