package route

import (
	"github.com/Daviskelvin824/OldEgg/middleware"
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/signup", controller.InsertUser)
	router.POST("/signin", controller.SignIn)
	router.POST("/authenticate", middleware.RequireAuthentication, controller.Authenticate)
	router.POST("/subscribe-to-newsletter", controller.SubscribeToNewsletter)
}
