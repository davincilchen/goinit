package usecase

const (
	GinKeyDevice = "Device"
)

type InfoCache interface {
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}

func AuthDeviceToken(cache InfoCache, token string) bool {

	dm := GetDeviceManager()
	dev := dm.GetByToken(token)
	if dev == nil {
		return false
	}
	CacheDevice(cache, dev)
	return true
}

func CacheDevice(ctx InfoCache, dev *LoginDevice) {
	ctx.Set(GinKeyDevice, dev)
}

func GetCacheDevice(ctx InfoCache) *LoginDevice {
	dev, exist := ctx.Get(GinKeyDevice)
	if !exist {
		return nil
	}
	if dev != nil {
		e, ok := dev.(*LoginDevice)
		if ok {
			return e
		}
	}
	return nil
}
