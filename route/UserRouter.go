package route

import (
	// "github.com/KevinChristian30/OldEgg/middleware"
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/signup", controller.InsertUser)
	router.POST("/signin", controller.SignIn)

}
