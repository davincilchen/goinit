package controllers

import (
	"encoding/json"
	"net/http"
	userUCase "xr-central/pkg/app/users/usecase"

	"github.com/gin-gonic/gin"
)

type loginSuccess func(*userUCase.LoginUser) error

type LoginController struct {
	ctx         *gin.Context
	loginParams UserLoginParams
	//loginUser   *userUCase.LoginUser
	fnSuccess loginSuccess
}

func (t *LoginController) Do() {

	ctx := t.ctx
	if t.loginParams.Account == nil || t.loginParams.Password == nil {
		RespBadRequest(ctx)
		return
	}

	loginUser := t.authWhenLogin()
	if loginUser == nil {
		RespInvalidPassword(ctx)
		return
	}

	err := t.fnSuccess(loginUser)

	if err != nil {
		RespUnknowError(ctx, err)
		return
	}
	LoginSucessReponse(ctx, loginUser)
}

func (t *LoginController) authWhenLogin() *userUCase.LoginUser {

	param := t.loginParams
	user := userUCase.User{}
	qRet, err := user.Login(*param.Account, *param.Password)
	if err != nil {
		return nil
	}

	return qRet

}

func NewLoginController(ctx *gin.Context,
	loginParams UserLoginParams,
	fn loginSuccess) *LoginController {
	return &LoginController{
		ctx:         ctx,
		loginParams: loginParams,
		fnSuccess:   fn,
	}
}

func getBodyForLogin(ctx *gin.Context, out interface{}) error {

	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	return err

}

func LoginSucessReponse(ctx *gin.Context, user *userUCase.LoginUser) {
	response := ResBody{}
	data := LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: user.Token,
	}
	response.ResCode = RES_OK
	response.Data = data
	ctx.JSON(http.StatusOK, response)

}
