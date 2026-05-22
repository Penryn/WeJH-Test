package userController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type bindForm struct {
	PassWord string `json:"password"`
}

func BindZFPassword(c *gin.Context) {
	var postForm bindForm
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
	if postForm.PassWord != "123456" {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func BindOauthPassword(c *gin.Context) {
	var postForm bindForm
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
	if postForm.PassWord != "123456" {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func BindLibraryPassword(c *gin.Context) {
	var postForm bindForm
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
	if postForm.PassWord != "123456" {
		_ = c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
