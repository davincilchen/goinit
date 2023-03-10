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

func (t *DeviceLoginProc) DevLoginSucess(ctx infopass.Context, user *userUCase.LoginUser) error {

	//TODO: save ip and login/logout
	device, err := deviceRepo.RegDevice(&t.Device)
	if err != nil {
		ctx.CacheDBError(err)
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
	//每次呼叫edge的ctx會不同,不能在new的時候跟著綁
	edge   *edgeUCase.Edge //not nil when post reserve success
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

func (t *LoginDevice) NewReserve(ctx infopass.Context, appID int) (*string, error) {
	if t.User == nil {
		return nil, errors.New("nil user for login device")
	}
	if t.IsReserve() {
		return nil, errDef.ErrRepeatedReserve
	}

	fmt.Println("appID", appID) //TODO remove

	manager := edgeUCase.GetEdgeManager()
	edge, err := manager.Reserve(ctx, appID)
	if err != nil {
		return nil, err
	}

	t.AttachEdge(edge)
	e := edge.GetInfo()
	return &e.IP, nil
}

func (t *LoginDevice) ReleaseReserve(ctx infopass.Context) error {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	if t.edge == nil {
		return nil
	}

	t.edge.ReleaseReserve(ctx)
	t.edge = nil
	return nil
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
