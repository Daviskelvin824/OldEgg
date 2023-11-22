package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func SearchQueryRoute(router *gin.Engine) {

	router.POST("/create-search-query", controller.CreateSearchQuery)
	router.GET("/get-popular-search-queries", controller.GetPopularSearchQueries)

}
