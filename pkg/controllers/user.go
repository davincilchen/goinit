package controllers

import (
	"net/http"
	userUCase "xr-central/pkg/app/users/usecase"

	"github.com/gin-gonic/gin"
)

type UserLoginParams struct {
	Account  *string
	Password *string
}

type LoginResponse struct { //:TODO for device login
	ID    uint   `json:"user_id"`
	Name  string `json:"user_name"`
	Token string
}

func Login(ctx *gin.Context) {
	req := &UserLoginParams{}
	err := getBodyForLogin(ctx, req)
	if err != nil {
		RespBadRequest(ctx)
		return
	}

	handle := NewLoginController(ctx, *req, UserLoginSucess)
	handle.Do()

}

func UserLoginSucess(user *userUCase.LoginUser) error {

	return nil
}

// func getBodyForUserLogin(ctx *gin.Context) (*UserLoginParams, error) {

// 	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var info UserLoginParams
// 	err = json.Unmarshal(body, &info)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &info, nil
// }

// func AuthWhenPlayerLogin(ctx *gin.Context) {
// 	param, err := getBodyForUserLogin(ctx)
// 	response := ResBody{}
// 	if err != nil {
// 		response.ResCode = RES_ERROR_BAD_REQUEST
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	userUCase := usecase.User{}
// 	qRet, err := userUCase.Login(*param.Account, *param.Password)
// 	if err != nil {
// 		response.ResCode = RES_INVALID_USER_PASSWORD
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	data := LoginResponse{
// 		ID:    qRet.ID,
// 		Name:  qRet.Name,
// 		Token: qRet.Token,
// 	}
// 	response.ResCode = RES_OK
// 	response.Data = data
// 	ctx.JSON(http.StatusOK, response)
// }

// .. //

func Logout(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}
