package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func OneTimeCodeRoute(router *gin.Engine) {

	router.POST("/get-one-time-code", controller.GetOneTimeCode)

}
