package api

import (
	"github.com/Zheaoli/nexus/controller"
	"github.com/gin-gonic/gin"
)

func APIInit(router *gin.Engine) {
	router.POST("/webhook/:source/:target/:url", controller.WebhookRequestController)

}
