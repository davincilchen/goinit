package deliverymodel

import (
	"xr-central/pkg/models"

	//devUCase "xr-central/pkg/app/device/usecase"
	edgeUCase "xr-central/pkg/app/edge/usecase"
)

type EdgeInfo struct {
	ID     uint                `json:"id"`
	IP     string              `json:"ip"`
	Port   int                 `json:"port"`
	Status models.EdgeStatus   `json:"status"`
	Online bool                `json:"online"`
	ActRet edgeUCase.ActionRet `json:"last_act_ret"`

	Device *DeviceInfo `json:"device"`
	AppID  *uint       `json:"app_id"`
}
