package usecase

import (
	"errors"
	"fmt"
	"sync"
	"xr-central/pkg/app/device/repo"
	"xr-central/pkg/app/infopass"
	"xr-central/pkg/models"

	errDef "xr-central/pkg/app/errordef"
	userUCase "xr-central/pkg/app/user/usecase"
)

var deviceRepo repo.Device

type DeviceLoginProc struct {
	Device    models.Device
	InfoCache infopass.InfoCache
}

func NewDeviceLoginProc(Type int, UUID string,
	InfoCache infopass.InfoCache) *DeviceLoginProc {
	d := &DeviceLoginProc{
		Device: models.Device{
			Type: Type,
			UUID: UUID,
		},
		InfoCache: InfoCache,
	}
	return d
}

func (t *DeviceLoginProc) DevLoginSucess(user *userUCase.LoginUser) error {

	//TODO: save ip and login/logout
	device, err := deviceRepo.RegDevice(&t.Device)
	if err != nil {
		infopass.CacheDBError(t.InfoCache, err)
		return err
	}

	loginDev := LoginDevice{
		Device: device,
		User:   user,
	}

	manager := GetDeviceManager()
	return manager.Add(&loginDev)

}

// ============================================= //
type LoginDevice struct {
	edgeMux sync.RWMutex
	edge    *models.Edge //not nil when post reserve
	Device  *models.Device
	User    *userUCase.LoginUser
}

func (t *LoginDevice) Logout() error {
	if t.User == nil {
		return errors.New("nil user for login device")
	}
	manager := GetDeviceManager()
	manager.Delete(t)
	return nil
}

func (t *LoginDevice) NewOrder(appID int) (*string, error) {
	if t.User == nil {
		return nil, errors.New("nil user for login device")
	}
	if t.IsReserve() {
		return nil, errDef.ErrRepeatedReserve
	}

	fmt.Println("appID", appID)

	// manager := GetDeviceManager()
	// manager.Delete(t)
	return nil, nil
}

func (t *LoginDevice) IsReserve() bool {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	return t.edge != nil

}
