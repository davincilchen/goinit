package controllers

import (
	"encoding/json"
	"net/http"
	"xr-central/pkg/app/users/usecase"

	"github.com/gin-gonic/gin"
)

func Deposit(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

func Withdraw(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

//users/0/deposit
func DepositTo(ctx *gin.Context) { //TODO:

}

func WithdrawFrom(ctx *gin.Context) { //TODO:

}

func Login(ctx *gin.Context) {
	AuthWhenPlayerLogin(ctx)
}

type LoginResponse struct {
	ID    uint
	Name  string
	Token string
}

type LoginParams struct {
	Name     string
	Password string
}

func getBodyForLogin(ctx *gin.Context) (*LoginParams, error) {

	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
	if err != nil {
		return nil, err
	}

	var info LoginParams
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func AuthWhenPlayerLogin(ctx *gin.Context) {
	param, err := getBodyForLogin(ctx)
	response := ResBody{}
	if err != nil {
		response.ResCode = RES_ERROR_BAD_REQUEST
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userUCase := usecase.User{}
	qRet, err := userUCase.Login(param.Name, param.Password)
	if err != nil {
		response.ResCode = RES_INVALID_USER_PASSWORD
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := LoginResponse{
		ID:    qRet.ID,
		Name:  qRet.Name,
		Token: qRet.Token,
	}
	response.ResCode = RES_OK
	response.Data = data
	ctx.JSON(http.StatusOK, response)
}

// .. //

func Logout(ctx *gin.Context) {

}
