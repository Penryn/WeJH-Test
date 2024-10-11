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

	result := []map[string]interface{}{
		{
			"className":   "心理健康与自我成长-0017",
			"credits":     "1.0",
			"lessonID":    "G207007",
			"lessonName":  "心理健康与自我成长",
			"score":       "100",
			"teacherName": "张苗苗",
		},
		{
			"className":   "高等数学Ⅰ-0012",
			"credits":     "5.0",
			"lessonID":    "G210013",
			"lessonName":  "高等数学Ⅰ",
			"score":       "100",
			"teacherName": "任博",
		},
		{
			"className":   "大学物理实验  A-0005",
			"credits":     "1.5",
			"lessonID":    "G410015",
			"lessonName":  "大学物理实验  A",
			"score":       "100",
			"teacherName": "张庆彬",
		},
		{
			"className":   "程序设计基础C-0005（大类）",
			"credits":     "4.0",
			"lessonID":    "G226002",
			"lessonName":  "程序设计基础C",
			"score":       "100",
			"teacherName": "毛国红",
		},
	}

	utils.JsonSuccessResponse(c, result)
}

func GetClassTable(c *gin.Context) {
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
		"info": map[string]interface{}{
			"ClassName": "2023计算机类05",
			"Name":      "wejh",
		},
		"lessonsTable": []map[string]interface{}{
			{
				"campus":      "朝晖校区",
				"classID":     "FC012572E79DE87BE0550113465EF1CF",
				"className":   "高等数学Ⅰ-0012",
				"credits":     "5.0",
				"id":          "G210013",
				"lessonHours": "80",
				"lessonName":  "高等数学Ⅰ",
				"lessonPlace": "教202",
				"placeID":     "19816",
				"sections":    "6-7",
				"teacherName": "任博",
				"type":        "必修课",
				"week":        "3-4周,7-19周",
				"weekday":     "1",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC40290D2471D2A4E0550113465EF1CF",
				"className":   "中国近现代史纲要-0006",
				"credits":     "2.0",
				"id":          "13082",
				"lessonHours": "32",
				"lessonName":  "中国近现代史纲要",
				"lessonPlace": "教401",
				"placeID":     "19829",
				"sections":    "8-9",
				"teacherName": "屈胜飞",
				"type":        "必修课",
				"week":        "3-4周,7-19周",
				"weekday":     "1",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC3EF90E6BAD9604E0550113465EF1CF",
				"className":   "线性代数B-0012",
				"credits":     "2.0",
				"id":          "BFD26E683632AC2EE0550113465EF1CF",
				"lessonHours": "32",
				"lessonName":  "线性代数B",
				"lessonPlace": "教202",
				"placeID":     "19816",
				"sections":    "1-2",
				"teacherName": "金建国",
				"type":        "必修课",
				"week":        "3-4周,6-19周",
				"weekday":     "2",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FE02D47D162B3774E0550113465EF1CF",
				"className":   "大学英语-0117",
				"credits":     "4.0",
				"id":          "4EAAB9FAF9641EB1E053A11310AC10C4",
				"lessonHours": "64",
				"lessonName":  "大学英语",
				"lessonPlace": "教404",
				"placeID":     "19832",
				"sections":    "3-4",
				"teacherName": "张其亮",
				"type":        "必修课",
				"week":        "3-4周,6-19周",
				"weekday":     "2",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC0690B0DAB78425E0550113465EF1CF",
				"className":   "程序设计基础C-0005（大类）",
				"credits":     "4.0",
				"id":          "4320",
				"lessonHours": "64",
				"lessonName":  "程序设计基础C",
				"lessonPlace": "教911",
				"placeID":     "E79443AB35B0D3F0E0550113465EF1CF",
				"sections":    "8-9",
				"teacherName": "毛国红",
				"type":        "必修课",
				"week":        "3周",
				"weekday":     "2",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC0690B0DAB78425E0550113465EF1CF",
				"className":   "程序设计基础C-0005（大类）",
				"credits":     "4.0",
				"id":          "4320",
				"lessonHours": "64",
				"lessonName":  "程序设计基础C",
				"lessonPlace": "教307",
				"placeID":     "19827",
				"sections":    "8-9",
				"teacherName": "毛国红",
				"type":        "必修课",
				"week":        "4-6周(双),7-19周",
				"weekday":     "2",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FAFE7696F8D0E4F0E0550113465EF1CF",
				"className":   "ACM程序设计竞赛入门-0001",
				"credits":     "3.0",
				"id":          "23139",
				"lessonHours": "48",
				"lessonName":  "ACM程序设计竞赛入门",
				"lessonPlace": "教911",
				"placeID":     "E79443AB35B0D3F0E0550113465EF1CF",
				"sections":    "10-12",
				"teacherName": "韩姗姗,侯向辉",
				"type":        "任选课",
				"week":        "3-4周,6-19周",
				"weekday":     "2",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC012572E79DE87BE0550113465EF1CF",
				"className":   "高等数学Ⅰ-0012",
				"credits":     "5.0",
				"id":          "G210013",
				"lessonHours": "80",
				"lessonName":  "高等数学Ⅰ",
				"lessonPlace": "子良A258",
				"placeID":     "19894",
				"sections":    "1-2",
				"teacherName": "任博",
				"type":        "必修课",
				"week":        "3-4周,6-19周",
				"weekday":     "3",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FBDECD37D3F4EED4E0550113465EF1CF",
				"className":   "大学物理实验  A-0005",
				"credits":     "1.5",
				"id":          "1059",
				"lessonHours": "48",
				"lessonName":  "大学物理实验  A",
				"lessonPlace": "实验室(朝晖)",
				"placeID":     "887F063496C84532E053A11310AC9FFF",
				"sections":    "3-5",
				"teacherName": "张庆彬",
				"type":        "必修课",
				"week":        "3-4周,6-19周",
				"weekday":     "3",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FE886A7E4EEDB3D2E0550113465EF1CF",
				"className":   "健美初级班男-陈春军周三67朝晖",
				"credits":     "1.0",
				"id":          "13861",
				"lessonHours": "32",
				"lessonName":  "体育",
				"lessonPlace": "游泳馆一楼力量房",
				"placeID":     "967FB36A95A41064E053A11310AC83BB",
				"sections":    "6-7",
				"teacherName": "陈春军",
				"type":        "体育课",
				"week":        "3-4周,6-19周",
				"weekday":     "3",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FE02D47D162B3774E0550113465EF1CF",
				"className":   "大学英语-0117",
				"credits":     "4.0",
				"id":          "4EAAB9FAF9641EB1E053A11310AC10C4",
				"lessonHours": "64",
				"lessonName":  "大学英语",
				"lessonPlace": "教404",
				"placeID":     "19832",
				"sections":    "3-4",
				"teacherName": "张其亮",
				"type":        "必修课",
				"week":        "3-4周,6-19周",
				"weekday":     "4",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC0690B0DAB78425E0550113465EF1CF",
				"className":   "程序设计基础C-0005（大类）",
				"credits":     "4.0",
				"id":          "4320",
				"lessonHours": "64",
				"lessonName":  "程序设计基础C",
				"lessonPlace": "教907",
				"placeID":     "E794BBAD394FD498E0550113465EF1CF",
				"sections":    "6-7",
				"teacherName": "毛国红",
				"type":        "必修课",
				"week":        "3-4周,6-11周",
				"weekday":     "4",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC0690B0DAB78425E0550113465EF1CF",
				"className":   "程序设计基础C-0005（大类）",
				"credits":     "4.0",
				"id":          "4320",
				"lessonHours": "64",
				"lessonName":  "程序设计基础C",
				"lessonPlace": "教307",
				"placeID":     "19827",
				"sections":    "6-7",
				"teacherName": "毛国红",
				"type":        "必修课",
				"week":        "12-19周",
				"weekday":     "4",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FBDE9052E632DA2DE0550113465EF1CF",
				"className":   "心理健康与自我成长-0017",
				"credits":     "1.0",
				"id":          "BFD4C9A5091AADA7E0550113465EF1CF",
				"lessonHours": "16",
				"lessonName":  "心理健康与自我成长",
				"lessonPlace": "教802",
				"placeID":     "19910",
				"sections":    "1-2",
				"teacherName": "张苗苗",
				"type":        "必修课",
				"week":        "3-4周,6-11周",
				"weekday":     "5",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FBF0F002BCE5200EE0550113465EF1CF",
				"className":   "心理健康教育实践-0017",
				"credits":     "1.0",
				"id":          "BFE5CE76EF65C404E0550113465EF1CF",
				"lessonHours": "4",
				"lessonName":  "心理健康教育实践",
				"lessonPlace": "教202",
				"placeID":     "19816",
				"sections":    "3-4",
				"teacherName": "张苗苗,孟婷婷",
				"type":        "必修课",
				"week":        "12周,19周",
				"weekday":     "5",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC012572E79DE87BE0550113465EF1CF",
				"className":   "高等数学Ⅰ-0012",
				"credits":     "5.0",
				"id":          "G210013",
				"lessonHours": "80",
				"lessonName":  "高等数学Ⅰ",
				"lessonPlace": "教201",
				"placeID":     "19815",
				"sections":    "8-9",
				"teacherName": "任博",
				"type":        "必修课",
				"week":        "3-4周,6-11周",
				"weekday":     "5",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "0589AE90ACBA7B7DE063A11310AC486A",
				"className":   "世界食品与文化欣赏-0003",
				"credits":     "2.0",
				"id":          "24053",
				"lessonHours": "32",
				"lessonName":  "世界食品与文化欣赏",
				"lessonPlace": "教401",
				"placeID":     "19829",
				"sections":    "10-11",
				"teacherName": "杜明明",
				"type":        "任选课",
				"week":        "3-4周,6-19周",
				"weekday":     "5",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC012572E79DE87BE0550113465EF1CF",
				"className":   "高等数学Ⅰ-0012",
				"credits":     "5.0",
				"id":          "G210013",
				"lessonHours": "80",
				"lessonName":  "高等数学Ⅰ",
				"lessonPlace": "教202",
				"placeID":     "19816",
				"sections":    "6-7",
				"teacherName": "任博",
				"type":        "必修课",
				"week":        "6周",
				"weekday":     "6",
			},
			{
				"campus":      "朝晖校区",
				"classID":     "FC40290D2471D2A4E0550113465EF1CF",
				"className":   "中国近现代史纲要-0006",
				"credits":     "2.0",
				"id":          "13082",
				"lessonHours": "32",
				"lessonName":  "中国近现代史纲要",
				"lessonPlace": "教401",
				"placeID":     "19829",
				"sections":    "8-9",
				"teacherName": "屈胜飞",
				"type":        "必修课",
				"week":        "6周",
				"weekday":     "6",
			},
		},
		"practiceLessons": []map[string]interface{}{
			{
				"className":   "4-11周",
				"credits":     "1.0",
				"lessonName":  "国家安全教育",
				"teacherName": "王绍让,顾容,蒋惠琴,闫丹,肖云泽,梁美赟,包士毅,李芸,吴杰,朱添田,刘杰,王婷",
			},
			{
				"className":   "1-16周",
				"credits":     "1.0",
				"lessonName":  "大学信息技术基础",
				"teacherName": "李强",
			},
		},
	}
	utils.JsonSuccessResponse(c, result)
}
