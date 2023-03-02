package usecase

import (
	"sync"
	errDef "xr-central/pkg/app/errordef"
)

type DeviceManager struct {
	deviceUUIDMap  map[string]*LoginDevice //KEY: UUID
	deviceTokenMap map[string]*LoginDevice //KEY: Token
	mux            sync.RWMutex
}

var deviceManager *DeviceManager

func newDeviceManager() *DeviceManager {
	d := &DeviceManager{}
	d.deviceUUIDMap = make(map[string]*LoginDevice)
	d.deviceTokenMap = make(map[string]*LoginDevice)

	return d
}

func GetDeviceManager() *DeviceManager {
	if deviceManager == nil {
		deviceManager = newDeviceManager()
	}
	return deviceManager
}

func (t *DeviceManager) Add(dev *LoginDevice) error {

	t.mux.Lock()
	defer t.mux.Unlock()

	_, ok := t.deviceUUIDMap[dev.Device.UUID]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}
	_, ok = t.deviceTokenMap[dev.User.Token]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}

	t.deviceUUIDMap[dev.Device.UUID] = dev
	t.deviceTokenMap[dev.User.Token] = dev
	return nil
}

func (t *DeviceManager) Get(uuid string) error {

	return nil
}
