package deliverymodel

import (
	devUCase "xr-central/pkg/app/device/usecase"
	// edgeUCase "xr-central/pkg/app/edge/usecase"
	// "xr-central/pkg/models"
)

type DeviceInfo struct {
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
		ID: devInfo.ID,
		IP: in.IP,
		User: UserInfo{
			ID:   userInfo.ID,
			Name: userInfo.Name,
		},
	}

}
