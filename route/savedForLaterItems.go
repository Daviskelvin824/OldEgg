package route

import (
	"github.com/Daviskelvin824/OldEgg/controller"
	"github.com/gin-gonic/gin"
)

func SavedForLaterItemsRoute(router *gin.Engine) {

	router.POST("/save-item-for-later", controller.SaveItemForLater)
	router.POST("/get-saved-for-later-items", controller.GetSavedForLaterItems)

}
