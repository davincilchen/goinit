package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// =========================================== //

type EdgeInfo struct {
	IP     string              `json:"ip"`
	Port   int                 `json:"port"`
	Status models.EdgeStatus   `json:"status"`
	Online bool                `json:"online"`
	ActRet edgeUCase.ActionRet `json:"last_act_ret"`
}

type EdgeStatusReq struct {
	DevStatus int    `json:"device_status"`
	StatusDes string `json:"status_des"`
}

type EdgeStatusResp struct {
	Edge *EdgeInfo `json:"edge"`
}

func edgeStatusParam(ctx *gin.Context) (*EdgeStatusReq, error) {
	// .. //
	param := EdgeStatusReq{}
	req, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		e := errors.New("read request body failed")
		return nil, e
	}
	err = json.Unmarshal(req, &param)
	if err != nil {
		e := errors.New("unmarshal body failed")
		return nil, e
	}
	return &param, nil
}

func EdgeStatus(ctx *gin.Context) {
	dev := devUCase.GetCacheDevice(ctx)
	if dev == nil {
		e := errors.New("GetCacheDevice Nil")
		dlv.RespError(ctx, e, nil)
		return
	}
	// .. //
	param, err := edgeStatusParam(ctx)
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}

	// .. //
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

	dev.UpdateStatus(ctxcache.NewContext(ctx), devUCase.DevStatus(param.DevStatus))
	ctx.JSON(http.StatusOK, response)
}

// =========================================== //

type EdgeListResp struct {
	Edges []EdgeInfo `json:"edges"`
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

// =========================================== //

type EdgeRegReq struct {
	AppsID []uint `json:"apps_id"`
	Port   int    `json:"port"`
}

func edgeRegParam(ctx *gin.Context) (*EdgeRegReq, error) {
	// .. //
	param := EdgeRegReq{}
	req, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		e := errors.New("read request body failed")
		return nil, e
	}
	err = json.Unmarshal(req, &param)
	if err != nil {
		e := errors.New("unmarshal body failed")
		return nil, e
	}

	if param.Port <= 0 {
		e := fmt.Errorf("error port: %d", param.Port)
		return nil, e
	}
	return &param, nil
}

func EdgeReg(ctx *gin.Context) {
	// .. //
	param, err := edgeRegParam(ctx)
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}
	//ip := ctx.Request.RemoteAddr
	ip := ctx.ClientIP() //TODO:127.0.0.1 ::1換成真的ip 否則dev會打不到
	manager := edgeUCase.GetEdgeManager()
	edge, err := manager.RegEdge(ip, param.Port)
	if err != nil {
		dlv.RespError(ctx, err, nil)
		return
	}
	edge.RegApps(param.AppsID)

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	//response.Data = data

	ctx.JSON(http.StatusOK, response)

}
