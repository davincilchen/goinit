package deilvery

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	errDef "xr-central/pkg/app/errordef"
	dlv "xr-central/pkg/delivery"

	devUCase "xr-central/pkg/app/device/usecase"
)

type NewOrderResp struct {
	GameServerIP string `json:"game_server_ip"`
}

func NewOrder(ctx *gin.Context) { //TODO:
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e)
		return
	}

	if dev.Edge != nil {
		dlv.RespError(ctx, errDef.ErrRepeatedReserve)
		return
	}

	data := NewOrderResp{}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func ReleaseOrder(ctx *gin.Context) { //TODO:
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}

func DeviceResume(ctx *gin.Context) { //TODO:
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_EDGE_LOST

	ctx.JSON(http.StatusOK, response)

}

type StartAppResp struct {
}

func StartApp(ctx *gin.Context) { //TODO:
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}

type StopAppResp struct {
}

func StopApp(ctx *gin.Context) { //TODO:
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

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

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
