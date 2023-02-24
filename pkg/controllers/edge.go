package controllers

import (

	//"xr-central/pkg/app/users/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

// func getBodyForLogin(ctx *gin.Context) (*LoginParams, error) {

// 	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var info LoginParams
// 	err = json.Unmarshal(body, &info)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &info, nil
// }

type NewOrderResp struct {
	GameServerIP string `json:"gameServerIP"`
}

func NewOrder(ctx *gin.Context) { //TODO:
	data := NewOrderResp{}

	response := ResBody{}
	response.ResCode = RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func ReleaseOrder(ctx *gin.Context) { //TODO:
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

func DeviceResume(ctx *gin.Context) { //TODO:
	type User struct {
		NameWantAcc string
	}

	user := User{NameWantAcc: "45"}
	ctx.JSON(200, user)
}

type StartAppResp struct {
}

func StartApp(ctx *gin.Context) { //TODO:
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

type StopAppResp struct {
}

func StopApp(ctx *gin.Context) { //TODO:

}

type EdgeStatusResp struct {
	Status int    `json:"status"`
	Online bool   `json:"online"`
	IP     string `json:"ip"`
}

func EdgeStatus(ctx *gin.Context) { //TODO:

	data := EdgeStatusResp{}

	response := ResBody{}
	response.ResCode = RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
