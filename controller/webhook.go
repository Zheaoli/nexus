package controller

import (
	"fmt"
	"github.com/Zheaoli/nexus/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlParser struct {
	Source string `uri:"source" `
	Target string `uri:"target" `
	URL    string `uri:"url" `
}

func WebhookRequestController(ctx *gin.Context) {
	var data UrlParser
	if err := ctx.ShouldBindUri(&data); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}
	err := handler.HandleWebHook(data.Source, data.Target, data.URL, ctx.Request.Body, ctx.Request.Header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	}
	return
}
