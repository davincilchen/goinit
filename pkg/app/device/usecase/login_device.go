package usecase

import (
	"errors"
	"fmt"
	"sync"
	repo "xr-central/pkg/app/device/repo/mysql"
	"xr-central/pkg/app/infopass"
	"xr-central/pkg/models"

	edgeUCase "xr-central/pkg/app/edge/usecase"
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
	edge    *edgeUCase.Edge //not nil when post reserve
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

	manager := edgeUCase.GetEdgeManager()
	edge, err := manager.Reserve(appID)
	if err != nil {
		return nil, err
	}

	t.AttachEdge(edge)

	return &edge.IP, nil
}

func (t *LoginDevice) IsReserve() bool {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	return t.edge != nil
}

func (t *LoginDevice) AttachEdge(edge *edgeUCase.Edge) {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	t.edge = edge
}
