package lostAndFoundRecordController

import (
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetKindList(c *gin.Context) {
	kinds := []map[string]interface{}{
		{
			"id":        1,
			"kind_name": "饭卡",
		},
		{
			"id":        2,
			"kind_name": "电子",
		},
		{
			"id":        3,
			"kind_name": "文体",
		},
		{
			"id":        4,
			"kind_name": "衣包",
		},
		{
			"id":        5,
			"kind_name": "证件",
		},
		{
			"id":        6,
			"kind_name": "其他",
		},
	}
	utils.JsonSuccessResponse(c, kinds)
}
