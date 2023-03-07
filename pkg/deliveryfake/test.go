package deliveryfake

import (
	"fmt"
	"net/http"

	"xr-central/pkg/delivery"

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

type NewReserveResp struct {
	GameServerIP string `json:"game_server_ip"`
}

func FakeNewReserve(ctx *gin.Context) { //TODO:
	fmt.Println("--------------")
	data := NewReserveResp{}

	response := delivery.ResBody{}
	response.ResCode = delivery.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func FakeReleaseReserve(ctx *gin.Context) { //TODO:
	response := delivery.ResBody{}
	response.ResCode = delivery.RES_OK

	ctx.JSON(http.StatusOK, response)
}

func FakeDeviceResume(ctx *gin.Context) { //TODO:
	response := delivery.ResBody{}
	response.ResCode = delivery.RES_EDGE_LOST

	ctx.JSON(http.StatusOK, response)

}

type StartAppResp struct {
}

func FakeStartApp(ctx *gin.Context) { //TODO:
	response := delivery.ResBody{}
	response.ResCode = delivery.RES_OK

	ctx.JSON(http.StatusOK, response)
}

type StopAppResp struct {
}

func FakeStopApp(ctx *gin.Context) { //TODO:
	response := delivery.ResBody{}
	response.ResCode = delivery.RES_OK

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

	response := delivery.ResBody{}
	response.ResCode = delivery.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
