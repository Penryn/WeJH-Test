package router

import (
	"wejh-go/app/controllers/funcControllers/lostAndFoundRecordController"
	"wejh-go/app/controllers/funcControllers/schoolBusController"
	"wejh-go/app/controllers/funcControllers/zfController"
	"wejh-go/app/controllers/yxyController/electricityController"
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
			zf.POST("/score", zfController.GetScore)
			zf.POST("/room", zfController.GetRoom)
		}
		lost := fun.Group("/lost", midwares.CheckLogin)
		{
			lost.GET("/kind_list", lostAndFoundRecordController.GetKindList)
		}
		electricity := fun.Group("/electricity", midwares.CheckLogin)
		{
			electricity.GET("/balance", electricityController.GetBalance)
		}
		bus := fun.Group("/bus", midwares.CheckLogin)
		{
			bus.GET("/list", schoolBusController.GetBusList)
		}
	}
}
