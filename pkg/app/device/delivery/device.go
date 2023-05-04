package delivery

import (
	"net/http"

	// "goinit/pkg/app/ctxcache"

	"github.com/gin-gonic/gin"

	dlvModel "goinit/pkg/app/deliverymodel"
	devUCase "goinit/pkg/app/device/usecase"
	dlv "goinit/pkg/delivery"
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
