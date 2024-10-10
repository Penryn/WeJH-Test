package zfController

import (
    "wejh-go/app/apiException"
    "wejh-go/app/services/sessionServices"
    "wejh-go/app/utils"

    "github.com/gin-gonic/gin"
)

type form struct {
    Year string `json:"year" binding:"required"`
    Term string `json:"term" binding:"required"`
}

func GetMidTermScore(c *gin.Context) {
    var postForm form
    err := c.ShouldBindJSON(&postForm)
    if err != nil {
        _ = c.AbortWithError(200, apiException.ParamError)
        return
    }
    _, err = sessionServices.GetUserSession(c)

    if err != nil {
        _ = c.AbortWithError(200, apiException.NotLogin)
        return
    }

    result := map[string]interface{}{
        "data": []map[string]interface{}{
            {
                "className":  "心理健康与自我成长-0017",
                "credits":    "1.0",
                "lessonID":   "G207007",
                "lessonName": "心理健康与自我成长",
                "score":      "100",
                "teacherName": "张苗苗",
            },
            {
                "className":  "高等数学Ⅰ-0012",
                "credits":    "5.0",
                "lessonID":   "G210013",
                "lessonName": "高等数学Ⅰ",
                "score":      "100",
                "teacherName": "任博",
            },
            {
                "className":  "大学物理实验  A-0005",
                "credits":    "1.5",
                "lessonID":   "G410015",
                "lessonName": "大学物理实验  A",
                "score":      "100",
                "teacherName": "张庆彬",
            },
            {
                "className":  "程序设计基础C-0005（大类）",
                "credits":    "4.0",
                "lessonID":   "G226002",
                "lessonName": "程序设计基础C",
                "score":      "100",
                "teacherName": "毛国红",
            },
        },
    }

    utils.JsonSuccessResponse(c, result)
}