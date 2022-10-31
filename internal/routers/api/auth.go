package api

import (
	"github.com/gin-gonic/gin"

	"github.com/mneumi/gin-boilerplate/global"
	"github.com/mneumi/gin-boilerplate/internal/service"
	"github.com/mneumi/gin-boilerplate/pkg/app"
	"github.com/mneumi/gin-boilerplate/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	err = svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
