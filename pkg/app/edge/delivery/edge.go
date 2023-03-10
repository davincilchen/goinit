package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"xr-central/pkg/app/ctxcache"
	devUCase "xr-central/pkg/app/device/usecase"
	edgeUCase "xr-central/pkg/app/edge/usecase"
	errDef "xr-central/pkg/app/errordef"
	dlv "xr-central/pkg/delivery"
)

type NewReserveResp struct {
	GameServerIP string `json:"game_server_ip"`
}

func NewReserve(ctx *gin.Context) { //TODO:
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		dlv.RespError(ctx, errDef.ErrUrlParamError, nil)
		return
	}
	nCtx := ctxcache.NewContext(ctx)
	ip, err := dev.NewReserve(nCtx, id)
	if err != nil || ip == nil {
		if err == errDef.ErrRepeatedReserve {
			dlv.RespError(ctx, errDef.ErrRepeatedReserve, nil)
		} else {
			dlv.RespError(ctx, errDef.ErrNoResource, err)
		}
		return
	}

	data := NewReserveResp{
		GameServerIP: *ip,
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func ReleaseReserve(ctx *gin.Context) { //TODO:
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}

	dev.ReleaseReserve(ctxcache.NewContext(ctx))
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
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status int    `json:"status"`
	Online bool   `json:"online"`
}

type EdgeListResp struct {
	Edge []EdgeStatusResp `json:"edge"`
}

func EdgeStatus(ctx *gin.Context) { //ODO:

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

func EdgeList(ctx *gin.Context) {
	manager := edgeUCase.GetEdgeManager()
	ret := manager.GetEdgeList()

	type Data struct {
		Edges []EdgeStatusResp `json:"edge_list"`
	}
	data := Data{}

	for _, v := range ret {
		tmp := EdgeStatusResp{
			IP:     v.IP,
			Port:   v.Port,
			Status: int(v.Status),
			Online: v.Online}
		data.Edges = append(data.Edges, tmp)

	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
