package delivery

import (
	"fmt"
	"net/http"
	devUCase "xr-central/pkg/app/device/usecase"
	userUCase "xr-central/pkg/app/user/usecase"
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
	Token string `json:"token"`
}

// ======================================== //

func DevLogin(ctx *gin.Context) {
	req := &DevLoginParams{}
	err := dlv.GetBodyFromRawData(ctx, req)
	if err != nil {
		dlv.RespBadRequest(ctx, err)
		return
	}

	if req.DevInfo.Type == nil || req.DevInfo.UUID == nil {
		e := fmt.Errorf("nil Type or UUID")
		dlv.RespBadRequest(ctx, e)
		return
	}

	d := devUCase.NewDeviceLoginProc(*req.DevInfo.Type, *req.DevInfo.UUID, ctx)
	handle := NewLoginController(ctx, req.UserLoginParams, d.DevLoginSucess)
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
		e := fmt.Errorf("nil Account or Password")
		dlv.RespBadRequest(ctx, e)
		return
	}

	loginUser := t.authWhenLogin()
	if loginUser == nil {
		dlv.RespInvalidPassword(ctx)
		return
	}

	err := t.fnSuccess(loginUser)

	if err != nil {
		dlv.RespUnknowError(ctx, err) //TODO:
		return
	}
	LoginSucessReponse(ctx, loginUser)
}

func (t *LoginController) authWhenLogin() *userUCase.LoginUser {

	param := t.loginParams
	user := userUCase.User{}
	ret, err := user.Login(*param.Account, *param.Password)
	if err != nil {
		return nil
	}

	return ret

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
		dlv.RespBadRequest(ctx, err)
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
