package schoolBusController

import (
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetBusList(c *gin.Context) {
    busList := []map[string]interface{}{
        {
            "id":          2,
            "line":        "直达线（朝晖-屏峰）",
            "departure":   "邵科馆",
            "destination": "语林楼",
            "type":        0,
            "startTime":   "20:30:00",
        },
        {
            "id":          3,
            "line":        "直达线（屏峰-朝晖）",
            "departure":   "语林楼",
            "destination": "邵科馆",
            "type":        1,
            "startTime":   "06:40:00",
        },
    }

    utils.JsonSuccessResponse(c, busList)
}
