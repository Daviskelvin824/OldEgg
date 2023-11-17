package route

import (
	// "github.com/KevinChristian30/OldEgg/middleware"
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/sign-up", controller.insertUser)
	// router.POST("/sign-in", controller.SignIn)
	// router.POST("/update-user", controller.UpdateUser)
	// router.POST("/authenticate", middleware.RequireAuthentication, controller.Authenticate)
	// router.POST("/get-users", controller.GetUsers)
	// router.POST("/get-one-time-sign-in-code", controller.GetOneTimeSignInCode)
	// router.POST("/sign-in-with-one-time-code", controller.SignInWithOneTimeCode)
	// router.POST("/request-forgot-password-code", controller.RequestForgotPasswordCode)
	// router.POST("/validate-forgot-password-code", controller.ValidateForgotPasswordCode)
	// router.POST("/reset-password", controller.ResetPassword)
	// router.POST("/request-two-factor-authentication-code", controller.RequestTFACode)
	// router.POST("/change-password", controller.ChangePassword)

	// router.POST("/subscribe-to-newsletter", controller.SubscribeToNewsletter)

	// router.GET("/get-visualization-data", controller.GetVisualizationData)

}
