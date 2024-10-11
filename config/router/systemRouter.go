package router

import (
	"wejh-go/app/controllers/systemController"

	"github.com/gin-gonic/gin"
)

func systemRouterInit(r *gin.RouterGroup) {
	r.POST("/applist", systemController.GetAppList)
}
