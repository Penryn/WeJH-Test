package router

import (
	"wejh-go/app/controllers/funcControllers/zfController"
	"wejh-go/app/midwares"

	"github.com/gin-gonic/gin"
)

func funcRouterInit(r *gin.RouterGroup) {
	fun := r.Group("/func")
	{
		zf := fun.Group("/zf", midwares.CheckLogin)
		{
			zf.POST("/midtermscore", zfController.GetMidTermScore)
			zf.POST("/classtable", zfController.GetClassTable)
			zf.POST("/exam", zfController.GetExam)
		}
	}
}
