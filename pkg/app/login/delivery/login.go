package delivery

import (
	"net/http"
	userUCase "xr-central/pkg/app/users/usecase"
	dlv "xr-central/pkg/delivery"

	"github.com/gin-gonic/gin"
)

type loginSuccess func(*userUCase.LoginUser) error

type UserLoginParams struct {
	Account  *string
	Password *string
}

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

type LoginResponse struct { //:TODO for device login
	ID    uint   `json:"user_id"`
	Name  string `json:"user_name"`
	Token string
}

// ======================================== //

func DevLogin(ctx *gin.Context) {
	req := &DevLoginParams{}
	err := dlv.GetBodyFromRawData(ctx, req)
	if err != nil {
		dlv.RespBadRequest(ctx)
		return
	}

	handle := NewLoginController(ctx, req.UserLoginParams, DevLoginSucess)
	handle.Do()

}

func DevLoginSucess(user *userUCase.LoginUser) error {

	return nil
}

// ======================================== //

type LoginController struct {
	ctx         *gin.Context
	loginParams UserLoginParams
	//loginUser   *userUCase.LoginUser
	fnSuccess loginSuccess
}

func (t *LoginController) Do() {

	ctx := t.ctx
	if t.loginParams.Account == nil || t.loginParams.Password == nil {
		dlv.RespBadRequest(ctx)
		return
	}

	loginUser := t.authWhenLogin()
	if loginUser == nil {
		dlv.RespInvalidPassword(ctx)
		return
	}

	err := t.fnSuccess(loginUser)

	if err != nil {
		dlv.RespUnknowError(ctx, err)
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

func LoginSucessReponse(ctx *gin.Context, user *userUCase.LoginUser) {
	response := dlv.ResBody{}
	data := LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: user.Token,
	}
	response.ResCode = dlv.RES_OK
	response.Data = data
	ctx.JSON(http.StatusOK, response)

}

// ======================================== //

func Login(ctx *gin.Context) {
	req := &UserLoginParams{}
	err := dlv.GetBodyFromRawData(ctx, req)
	if err != nil {
		dlv.RespBadRequest(ctx)
		return
	}

	handle := NewLoginController(ctx, *req, UserLoginSucess)
	handle.Do()

}

func UserLoginSucess(user *userUCase.LoginUser) error {

	return nil
}

// ======================================== //

func Logout(ctx *gin.Context) {
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}