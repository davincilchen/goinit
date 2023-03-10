package usecase

import (
	"fmt"
	"sync"
	"xr-central/pkg/models"

	edgeHttp "xr-central/pkg/app/edge/repo/http"
	errDef "xr-central/pkg/app/errordef"
	"xr-central/pkg/app/infopass"
)

type HttpEdge interface {
	SetURL(url string)
	Reserve(ctx infopass.Context, appID int) error
	Release(ctx infopass.Context) error
	Resume(ctx infopass.Context) error
	StartAPP(ctx infopass.Context, appID int) error
	StopAPP(ctx infopass.Context) error
	GetStatus(ctx infopass.Context) error
}

type Edge struct {
	mux   sync.Mutex
	info  models.Edge
	eHttp HttpEdge
}

func NewEdge(edge models.Edge) *Edge {
	e := Edge{
		info: edge,
	}
	e.eHttp = edgeHttp.NewEdge(e.GetURL())
	return &e
}

func (t *Edge) GetURL() string {
	e := t.info
	if e.Port > 0 {
		return fmt.Sprintf("%s:%d", e.IP, e.Port)
	}

	return e.IP
}

func (t *Edge) Reserve(ctx infopass.Context, appID int) error {

	//online由每次reg時確認,減少api時間
	ok := t.updateStatusWhen(models.STATUS_FREE, models.STATUS_RESERVE_INIT)
	if !ok {
		return errDef.ErrNoResource
	}

	err := t.eHttp.Reserve(ctx, appID)

	status := models.STATUS_RESERVE_XR_NOT_CONNECT
	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	}

	t.updateStatus(status, online)
	return err

}

func (t *Edge) ReleaseReserve(ctx infopass.Context) error {

	status, _ := t.GetCacheStatus()
	if status == models.STATUS_FREE {
		return errDef.ErrAlreadyFree
	}

	status = models.STATUS_RX_RELEASE
	t.updateStatus(status, nil)
	err := t.eHttp.Release(ctx)

	status = models.STATUS_FREE
	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	}

	t.updateStatus(status, online)
	return err

}

func (t *Edge) Resume(ctx infopass.Context) error {

	err := t.eHttp.Resume(ctx)

	if err == errDef.ErrEdgeLost {
		t.setOnline(false)
	}

	return err
}

func (t *Edge) GetStatus() error {

	return nil
}

func (t *Edge) GetCacheStatus() (models.EdgeStatus, bool) {
	t.mux.Lock()
	defer t.mux.Unlock()

	e := t.info
	return e.Status, e.Online
}

func (t *Edge) StartAPP(ctx infopass.Context, appID int) error {

	ok := t.updateStatusWhen(models.STATUS_RESERVE_XR_CONNECT,
		models.STATUS_RX_START_APP)
	if !ok {
		return errDef.ErrCloudXRUnconect
	}

	err := t.eHttp.StartAPP(ctx, appID)

	status := models.STATUS_PLAYING
	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	}

	t.updateStatus(status, online)
	return err
}

func (t *Edge) StopAPP(ctx infopass.Context) error {

	ok := t.updateStatusWhen(models.STATUS_PLAYING,
		models.STATUS_RX_STOP_APP)
	if !ok {
		return errDef.ErrNotPlaying
	}

	err := t.eHttp.StopAPP(ctx)

	status := models.STATUS_RESERVE_XR_CONNECT
	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	}

	t.updateStatus(status, online)

	return err
}

func (t *Edge) OnXRConnect() {
	online := true
	t.updateStatus(models.STATUS_RESERVE_XR_CONNECT, &online)
	//updateStatus when
}

func (t *Edge) updateStatus(status models.EdgeStatus, online *bool) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.info.Status = status

	if online != nil {
		t.info.Online = *online
	}
}

func (t *Edge) updateStatusWhen(oriStatus, newStatus models.EdgeStatus) bool {
	t.mux.Lock()
	defer t.mux.Unlock()

	if !t.info.Online {
		return false
	}

	if t.info.Status == oriStatus {
		t.info.Status = newStatus
		return true
	}

	return false
}

func (t *Edge) setOnline(online bool) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.info.Online = online
}

func (t *Edge) GetInfo() models.Edge { //副本
	return t.info
}
