package usecase

import (
	"sync"
	"time"
	errDef "xr-central/pkg/app/errordef"

	cache "github.com/patrickmn/go-cache"
)

type DeviceManager struct {
	deviceUUIDMap      map[string]*LoginDevice //KEY: UUID
	deviceTokenMap     map[string]*LoginDevice //KEY: Token
	edgeIDtoDevUUIDMap map[uint]string         //KEY: edgeID
	mux                sync.RWMutex

	uuidCache *cache.Cache
}

var deviceManager *DeviceManager
var dafaultKeepAliveInterval time.Duration = 5 * time.Minute

func newDeviceManager() *DeviceManager {
	d := &DeviceManager{}
	d.deviceUUIDMap = make(map[string]*LoginDevice)
	d.deviceTokenMap = make(map[string]*LoginDevice)
	d.edgeIDtoDevUUIDMap = make(map[uint]string)
	return d
}

func GetDeviceManager() *DeviceManager {
	if deviceManager == nil {
		deviceManager = newDeviceManager()
		deviceManager.uuidCache = cache.New(dafaultKeepAliveInterval,
			10*time.Minute)

	}
	return deviceManager
}

func (t *DeviceManager) Add(dev *LoginDevice) error {

	t.mux.Lock()
	defer func() {
		t.mux.Unlock()
		t.uuidCache.Set(dev.device.UUID, dev.device.UUID, cache.DefaultExpiration)
	}()

	_, ok := t.deviceUUIDMap[dev.device.UUID]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}
	_, ok = t.deviceTokenMap[dev.user.Token]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}

	t.deviceUUIDMap[dev.device.UUID] = dev
	t.deviceTokenMap[dev.user.Token] = dev
	return nil
}

func (t *DeviceManager) Alive(uuid string) {
	t.uuidCache.Set(uuid, uuid, cache.DefaultExpiration)
}

func (t *DeviceManager) reserveFor(edgeID uint, devUUID string) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.edgeIDtoDevUUIDMap[edgeID] = devUUID

	return nil
}

func (t *DeviceManager) releseReserve(edgeID uint) error {
	t.mux.Lock()
	defer t.mux.Unlock()

	delete(t.edgeIDtoDevUUIDMap, edgeID)
	return nil
}

func (t *DeviceManager) GetByUUID(token string) *LoginDevice {

	t.mux.RLock()
	defer t.mux.RUnlock()

	dev, ok := t.deviceUUIDMap[token]
	if ok {
		return dev
	}
	return nil
}

func (t *DeviceManager) GetByToken(uuid string) *LoginDevice {

	t.mux.RLock()
	defer t.mux.RUnlock()

	dev, ok := t.deviceTokenMap[uuid]
	if ok {
		return dev
	}
	return nil
}

func (t *DeviceManager) Delete(dev *LoginDevice) {

	t.mux.Lock()
	defer t.mux.Unlock()

	delete(t.deviceTokenMap, dev.user.Token)
	delete(t.deviceUUIDMap, dev.device.UUID)

	if dev.edge != nil {
		delete(t.edgeIDtoDevUUIDMap, dev.edge.GetInfo().ID)
	}
}

func (t *DeviceManager) GetDevInfoWithEdge(edgeID uint) *QLoginDeviceRet {
	t.mux.Lock()
	defer t.mux.Unlock()

	uuid, ok := t.edgeIDtoDevUUIDMap[edgeID]
	if !ok {
		return nil
	}

	dev, ok := t.deviceUUIDMap[uuid]
	if !ok {
		return nil
	}

	ret := dev.GetDeviceInfo()
	return &ret

}
