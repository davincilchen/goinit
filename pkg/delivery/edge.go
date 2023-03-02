package delivery

import (

	//"xr-central/pkg/app/user/usecase"

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
	GameServerIP string `json:"game_server_ip"`
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
	response := ResBody{}
	response.ResCode = RES_EDGE_LOST

	ctx.JSON(http.StatusOK, response)

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
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

type EdgeStatusResp struct {
	Status int    `json:"status"`
	Online bool   `json:"online"`
	IP     string `json:"ip"`
}

func EdgeStatus(ctx *gin.Context) { //TODO:

	type Data struct {
		Edge EdgeStatusResp `json:"edge"`
	}
	data := Data{}
	data.Edge = EdgeStatusResp{}

	response := ResBody{}
	response.ResCode = RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
