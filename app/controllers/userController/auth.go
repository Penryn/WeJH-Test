package userController

import (
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/config/wechat"

	"github.com/gin-gonic/gin"
)

type autoLoginForm struct {
	Code      string `json:"code" binding:"required"`
	LoginType string `json:"type"`
}
type passwordLoginForm struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	LoginType string `json:"type"`
}


var user models.User


func AuthByPassword(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)

	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user = models.User{
		ID:           1,
		Username:     "wejh",
		JHPassword:     "123456",
		StudentID:    "202103150901",
		PhoneNum:     "12345678901",
		Type: 3,
		LibPassword: "123456",
		ZFPassword: "123456",
		OauthPassword: "123456",
		YxyUid: "123456",
		UnionID: "123456",
		CreateTime: time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),

	}

	err = sessionServices.SetUserSession(c, &user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf":    user.ZFPassword != "",
				"lib":   user.LibPassword != "",
				"yxy":   user.YxyUid != "",
				"oauth": user.OauthPassword != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})

}

func AuthBySession(c *gin.Context) {
	user, err := sessionServices.UpdateUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf":  user.ZFPassword != "",
				"lib": user.LibPassword != "",
				"yxy": user.YxyUid != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}

func WeChatLogin(c *gin.Context) {
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		_ = c.AbortWithError(200, apiException.OpenIDError)
		return
	}
	var user models.User
	if session.OpenID != "" {
		user = models.User{
			ID:           1,
			Username:     "wejh",
			JHPassword:     "123456",
			StudentID:    "202103150901",
			PhoneNum:     "12345678901",
			Type: 3,
			LibPassword: "123456",
			ZFPassword: "123456",
			OauthPassword: "123456",
			YxyUid: "123456",
			UnionID: "123456",
			CreateTime: time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),
	
		}
	} else {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}
	err = sessionServices.SetUserSession(c, &user)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf":    user.ZFPassword != "",
				"lib":   user.LibPassword != "",
				"yxy":   user.YxyUid != "",
				"oauth": user.OauthPassword != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}
