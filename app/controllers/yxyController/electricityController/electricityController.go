package electricityController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	_, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	electricityBalance := map[string]interface{}{
		"area_id": "1908211437059099",
        "building_code": "54",
        "display_room_name": "屏峰校区家和西苑13号楼3层x13306",
        "floor_code": "5403",
        "md_name": "照明用电",
        "md_type": "53220",
        "room_code": "540306",
        "room_status": "正常用电",
        "school_code": "10337",
        "soc": 158.05,
        "soc_amount": 85.03,
        "subsidy": 0,
        "subsidy_amount": 0,
        "surplus": 158.05,
        "surplus_amount": 85.03,
	}
	utils.JsonSuccessResponse(c, electricityBalance)
}
