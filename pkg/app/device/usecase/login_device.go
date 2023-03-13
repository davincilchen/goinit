package usecase

import (
	"errors"
	"sync"
	"xr-central/pkg/app/ctxcache"
	repo "xr-central/pkg/app/device/repo/mysql"
	"xr-central/pkg/models"

	edgeUCase "xr-central/pkg/app/edge/usecase"
	errDef "xr-central/pkg/app/errordef"
	userUCase "xr-central/pkg/app/user/usecase"
)

type DevStatus int

const (
	STATUS_FREE DevStatus = 0
	//STATUS_RESERVE_INIT           DevStatus = 110
	STATUS_RESERVE_XR_NOT_CONNECT DevStatus = 120
	STATUS_RESERVE_XR_CONNECT     DevStatus = 130
	//STATUS_RX_START_APP           DevStatus = 140
	STATUS_PLAYING DevStatus = 150
	// STATUS_RX_STOP_APP            DevStatus = 160
	// STATUS_RX_RELEASE             DevStatus = 170
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

func (t *DeviceLoginProc) DevLoginSucess(ctx ctxcache.Context, user *userUCase.LoginUser) error {

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

	statusMux sync.RWMutex
	status    DevStatus

	appMux sync.RWMutex
	appID  int
}

func (t *LoginDevice) Logout(ctx ctxcache.Context) error {
	if t.User == nil {
		return errors.New("nil user for login device")
	}
	_ = t.ReleaseReserve(ctx)
	manager := GetDeviceManager()
	manager.Delete(t)
	return nil
}

func (t *LoginDevice) NewReserve(ctx ctxcache.Context, appID int) (*string, error) {
	if t.User == nil {
		return nil, errors.New("nil user for login device")
	}
	if t.IsReserve() {
		return nil, errDef.ErrRepeatedReserve
	}

	manager := edgeUCase.GetEdgeManager()
	edge, err := manager.Reserve(ctx, appID)
	if err != nil {
		return nil, err
	}

	t.AttachEdge(edge)
	e := edge.GetInfo()
	t.SetAppID(appID)
	return &e.IP, nil
}

func (t *LoginDevice) ReleaseReserve(ctx ctxcache.Context) error {
	t.statusMux.Lock()
	t.status = STATUS_FREE
	t.statusMux.Unlock()

	edge := t.getEdge()
	if edge == nil {
		return errDef.ErrDevNoReserve
	}
	edge.ReleaseReserve(ctx)
	t.DetachEdge()

	t.SetAppID(0)
	return nil
}

func (t *LoginDevice) StartApp(ctx ctxcache.Context) error {
	edge := t.getEdge()
	if edge == nil {
		return errDef.ErrDevNoReserve
	}
	appID := t.GetAppID()
	return edge.StartAPP(ctx, appID)
}

func (t *LoginDevice) StopApp(ctx ctxcache.Context) error {
	edge := t.getEdge()
	if edge == nil {
		return errDef.ErrDevNoReserve
	}
	return edge.StopAPP(ctx)
}

func (t *LoginDevice) Resume(ctx ctxcache.Context) error {
	edge := t.getEdge()
	if edge == nil {
		return errDef.ErrDevNoReserve
	}
	return edge.Resume(ctx)
}

func (t *LoginDevice) OnCloudXRConnect(ctx ctxcache.Context) error {
	edge := t.getEdge()
	if edge == nil {
		return errDef.ErrDevNoReserve
	}
	return edge.OnCloudXRConnect(ctx)
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

func (t *LoginDevice) DetachEdge() {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	t.edge = nil
}

func (t *LoginDevice) getEdge() *edgeUCase.Edge {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()
	if t.edge == nil {
		return nil
	}
	return t.edge
}

func (t *LoginDevice) GetEdgeInfo() *edgeUCase.EdgeInfoStatus {
	t.edgeMux.Lock()
	defer t.edgeMux.Unlock()

	if t.edge == nil {
		return nil
	}

	e := t.edge.GetInfo()
	return &e
}

func (t *LoginDevice) GetAppID() int {
	t.appMux.RLock()
	defer t.appMux.RUnlock()
	return t.appID
}

func (t *LoginDevice) SetAppID(appID int) {
	t.appMux.Lock()
	defer t.appMux.Unlock()
	t.appID = appID
}
