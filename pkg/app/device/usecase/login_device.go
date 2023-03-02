package usecase

import (
	"fmt"
	"xr-central/pkg/app/device/repo"
	userUCase "xr-central/pkg/app/user/usecase"
	"xr-central/pkg/models"
)

var deviceRepo repo.Device

type DeviceLoginProc struct {
	Device models.Device
}

func NewDeviceLoginProc(Type int, UUID string) *DeviceLoginProc {
	d := &DeviceLoginProc{
		Device: models.Device{
			Type: Type,
			UUID: UUID,
		},
	}
	return d
}

func (t *DeviceLoginProc) DevLoginSucess(user *userUCase.LoginUser) error {

	fmt.Printf("%+v\n", t)

	device, err := deviceRepo.RegDevice(&t.Device)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", device)
	return nil
}
