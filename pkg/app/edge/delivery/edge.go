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
	"xr-central/pkg/models"
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

	err := dev.ReleaseReserve(ctxcache.NewContext(ctx))
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}

func DeviceResume(ctx *gin.Context) { //TODO:
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}

	err := dev.Resume(ctxcache.NewContext(ctx))
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)

}

type StartAppResp struct {
}

func StartApp(ctx *gin.Context) { //TODO:
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}

	err := dev.StartApp(ctxcache.NewContext(ctx))
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}

type StopAppResp struct {
}

func StopApp(ctx *gin.Context) {
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}

	err := dev.StopApp(ctxcache.NewContext(ctx))
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK

	ctx.JSON(http.StatusOK, response)
}

type EdgeInfo struct {
	IP     string              `json:"ip"`
	Port   int                 `json:"port"`
	Status models.EdgeStatus   `json:"status"`
	Online bool                `json:"online"`
	ActRet edgeUCase.ActionRet `json:"last_act_ret"`
}

type EdgeStatusResp struct {
	Edge *EdgeInfo `json:"edge"`
}

type EdgeListResp struct {
	Edges []EdgeInfo `json:"edges"`
}

func EdgeStatus(ctx *gin.Context) {

	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}
	edge := dev.GetEdgeInfo()

	data := EdgeStatusResp{}
	if edge != nil {
		tmp := EdgeInfo{
			IP:     edge.IP,
			Port:   edge.Port,
			Status: edge.Status,
			Online: edge.Online,
			ActRet: edge.ActRet,
		}
		data.Edge = &tmp
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}

func EdgeList(ctx *gin.Context) {
	manager := edgeUCase.GetEdgeManager()
	ret := manager.GetEdgeList()

	data := EdgeListResp{}

	for _, v := range ret {
		tmp := EdgeInfo{
			IP:     v.IP,
			Port:   v.Port,
			Status: v.Status,
			Online: v.Online,
			ActRet: v.ActRet,
		}
		data.Edges = append(data.Edges, tmp)

	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
