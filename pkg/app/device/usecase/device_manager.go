package usecase

import (
	"fmt"
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

//var dafaultKeepAliveInterval time.Duration = 5 * time.Minute
//var dafaultCleanAliveInterval time.Duration = 5 * time.Second

// var dafaultKeepAliveInterval time.Duration = 1 * time.Minute
// var dafaultCleanAliveInterval time.Duration = 1 * time.Second

var dafaultKeepAliveInterval time.Duration = 10 * time.Second
var dafaultCleanAliveInterval time.Duration = 11 * time.Second

func newDeviceManager() *DeviceManager {
	d := &DeviceManager{}
	d.deviceUUIDMap = make(map[string]*LoginDevice)
	d.deviceTokenMap = make(map[string]*LoginDevice)
	d.edgeIDtoDevUUIDMap = make(map[uint]string)
	d.uuidCache = cache.New(dafaultKeepAliveInterval,
		dafaultCleanAliveInterval)
	d.uuidCache.OnEvicted(d.reserveTimeout)
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

	_, ok := t.deviceTokenMap[dev.user.Token]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}

	_, ok = t.deviceUUIDMap[dev.device.UUID]
	if ok {
		return errDef.ErrRepeatedLogin //請先登出
	}

	t.deviceUUIDMap[dev.device.UUID] = dev
	t.deviceTokenMap[dev.user.Token] = dev
	return nil
}

func (t *DeviceManager) Alive(uuid string) {
	_, ok := t.uuidCache.Get(uuid)
	if !ok {
		return
	}
	t.uuidCache.Set(uuid, uuid, cache.DefaultExpiration)
}

func (t *DeviceManager) reserveTimeout(uuid string, value interface{}) {

	edgeID := uint(0)
	edgeIP := ""
	devID := uint(0)
	t.mux.Lock()
	dev, ok := t.deviceUUIDMap[uuid]
	t.mux.Unlock()

	if ok && dev != nil {
		devID = dev.device.ID
		edge := dev.GetEdgeInfo()
		if edge != nil {
			edgeID = edge.ID
			edgeIP = edge.IP
		}
	}
	fmt.Println(time.Now(), " [ReserveTimeout] edge_id:", edgeID,
		",IP:", edgeIP,
		",dev_id:", devID,
		", ", uuid)
}

func (t *DeviceManager) reserveFor(edgeID uint, devUUID string) error {

	t.uuidCache.Set(devUUID, devUUID, cache.DefaultExpiration)

	t.mux.Lock()
	defer t.mux.Unlock()
	t.edgeIDtoDevUUIDMap[edgeID] = devUUID

	return nil
}

func (t *DeviceManager) releseReserve(edgeID uint, uuid string) error {

	t.uuidCache.Delete(uuid)

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
