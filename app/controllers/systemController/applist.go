package systemController

import (
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetAppList(c *gin.Context) {
    appLists := []map[string]interface{}{
        {
            "id":              1,
            "title":           "课表查询",
            "route":           "/pages/lessonstable/index",
            "backgroundColor": "blue",
            "icon":            "https://gw.alicdn.com/tfs/TB1zXRm4uL2gK0jSZPhXXahvXXa-652-652.png",
            "require":         "zf",
        },
        {
            "id":              2,
            "title":           "成绩查询",
            "route":           "/pages/score/index",
            "backgroundColor": "green",
            "icon":            "https://gw.alicdn.com/tfs/TB1ZG8t4EY1gK0jSZFCXXcwqXXa-652-652.png",
            "require":         "zf",
        },
        {
            "id":              3,
            "title":           "考试安排",
            "route":           "/pages/exam/index",
            "backgroundColor": "orange",
            "icon":            "https://gw.alicdn.com/tfs/TB1oPcIrSR26e4jSZFEXXbwuXXa-652-652.png",
            "require":         "zf",
        },
        {
            "id":              4,
            "title":           "空教室",
            "route":           "/pages/freeroom/index",
            "backgroundColor": "green",
            "icon":            "https://gw.alicdn.com/tfs/TB13t8D4ET1gK0jSZFrXXcNCXXa-652-652.png",
            "require":         "zf",
        },
        {
            "id":              5,
            "title":           "校园卡",
            "route":           "/pages/schoolcard/index",
            "backgroundColor": "green",
            "icon":            "https://gw.alicdn.com/tfs/TB10a8t4EY1gK0jSZFCXXcwqXXa-652-652.png",
            "require":         "yxy",
        },
        {
            "id":              6,
            "title":           "借阅信息",
            "route":           "/pages/library/index",
            "backgroundColor": "yellow",
            "icon":            "https://gw.alicdn.com/tfs/TB1bYcjpODsXe8jSZR0XXXK6FXa-652-652.png",
            "require":         "library",
        },
        {
            "id":              7,
            "title":           "电费查询",
            "route":           "/pages/electricity/index",
            "backgroundColor": "green",
            "icon":            "https://img.cnpatrickstar.com/a6c9a873-e7c4-427a-8408-69b5c8ba887e.webp",
            "require":         "yxy",
        },
        {
            "id":              8,
            "title":           "校车信息",
            "route":           "/pages/schoolbus/index",
            "backgroundColor": "green",
            "icon":            "https://gw.alicdn.com/tfs/TB1ZG8t4EY1gK0jSZFCXXcwqXXa-652-652.png",
            "require":         "null",
        },
        {
            "id":              9,
            "title":           "正装借用",
            "route":           "/pages/suit/index",
            "backgroundColor": "orange",
            "icon":            "https://gw.alicdn.com/tfs/TB1zXRm4uL2gK0jSZPhXXahvXXa-652-652.png",
            "require":         "null",
        },
    }

    utils.JsonSuccessResponse(c, appLists)
}
