package usecase

import (
	"errors"
	"xr-central/pkg/app/device/repo"
	"xr-central/pkg/app/infopass"
	userUCase "xr-central/pkg/app/user/usecase"
	"xr-central/pkg/models"
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
	Edge   *models.Edge //not nil when post reserve
	Device *models.Device
	User   *userUCase.LoginUser
}

func (t *LoginDevice) Logout() error {
	if t.User == nil {
		return errors.New("nil user for login device")
	}
	manager := GetDeviceManager()
	manager.Delete(t)
	return nil
}
