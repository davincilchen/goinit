package deliverymodel

import (
	devUCase "xr-central/pkg/app/device/usecase"
	// edgeUCase "xr-central/pkg/app/edge/usecase"
	// "xr-central/pkg/models"
)

type DeviceInfoDetail struct {
	DeviceInfo
	Edge *EdgeInfo `json:"edge"`
}

type DeviceInfo struct {
	UUID string   `json:"uuid"`
	ID   uint     `json:"id"`
	IP   string   `json:"ip"`
	User UserInfo `json:"user"`
}

func WarpDeviceInfo(in *devUCase.QLoginDeviceRet) *DeviceInfo {

	if in == nil {
		return nil
	}
	devInfo := in.Device
	userInfo := in.User
	return &DeviceInfo{
		UUID: devInfo.UUID,
		ID:   devInfo.ID,
		IP:   in.IP,
		User: UserInfo{
			ID:   userInfo.ID,
			Name: userInfo.Name,
		},
	}

}