package delivery

import (
	"net/http"

	// "xr-central/pkg/app/ctxcache"

	"github.com/gin-gonic/gin"

	dlvModel "xr-central/pkg/app/deliverymodel"
	devUCase "xr-central/pkg/app/device/usecase"
	dlv "xr-central/pkg/delivery"
)

type DeviceDetailListResp struct {
	Total int                         `json:"total_num"`
	List  []dlvModel.DeviceInfoDetail `json:"device"`
}

func DeviceList(ctx *gin.Context) {

	devM := devUCase.GetDeviceManager()
	list := make([]dlvModel.DeviceInfoDetail, 0)

	devices := devM.GetDevices()
	for _, v := range devices {
		warpDev := dlvModel.WarpDeviceInfo(&v.QLoginDeviceRet)
		tmp := dlvModel.DeviceInfoDetail{
			DeviceInfo: *warpDev,
			Edge:       dlvModel.WarpEdgeInfo(v.Edge, &v.QLoginDeviceRet),
		}
		list = append(list, tmp)
	}

	data := DeviceDetailListResp{
		Total: len(list),
		List:  list,
	}
	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
