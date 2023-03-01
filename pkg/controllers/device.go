package controllers

import (
	userUCase "xr-central/pkg/app/users/usecase"

	"github.com/gin-gonic/gin"
)

type DevInfo struct {
	Type *int
	UUID *string
}

type DevLoginParams struct {
	UserLoginParams
	DevInfo
}

type DevLoginResponse struct {
	ID    uint
	Name  string
	Token string
}

func DevLogin(ctx *gin.Context) {
	req := &DevLoginParams{}
	err := getBodyForLogin(ctx, req)
	if err != nil {
		RespBadRequest(ctx)
		return
	}

	handle := NewLoginController(ctx, req.UserLoginParams, DevLoginSucess)
	handle.Do()

}

func DevLoginSucess(user *userUCase.LoginUser) error {

	return nil
}
