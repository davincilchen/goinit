package usecase

import (
	"sync"
	"xr-central/pkg/models"

	edgeHttp "xr-central/pkg/app/edge/repo/http"
	errDef "xr-central/pkg/app/errordef"
)

type HttpEdge interface {
	Release() error
	Resume() error
	StartAPP() error
	StopAPP() error
	GetStatus() error
}

type Edge struct {
	mux sync.RWMutex
	models.Edge

	eHttp HttpEdge
}

func NewEdge() *Edge {
	e := Edge{
		eHttp: &edgeHttp.Edge{},
	}
	return &e
}

func (t *Edge) Release() error {

	status, _ := t.GetCacheStatus()
	if status == models.STATUS_FREE {
		return errDef.ErrAlreadyFree
	}

	status = models.STATUS_RX_RELEASE
	t.updateStatus(status, nil)
	err := t.eHttp.Release()

	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	} else {
		status = models.STATUS_FREE
	}

	t.updateStatus(status, online)
	return err

}

func (t *Edge) Resume() error {

	err := t.eHttp.Resume()

	if err == errDef.ErrEdgeLost {
		t.setOnline(false)
	}

	return err
}

func (t *Edge) GetStatus() error {

	return nil
}

func (t *Edge) GetCacheStatus() (models.EdgeStatus, bool) {
	t.mux.RLock()
	defer t.mux.RUnlock()

	return t.Status, t.Online
}

func (t *Edge) StartAPP() error {
	status, _ := t.GetCacheStatus()
	if status != models.STATUS_RESERVE_XR_CONNECT {
		return errDef.ErrCloudXRUnconect
	}

	status = models.STATUS_RX_START_APP
	t.updateStatus(status, nil)

	err := t.eHttp.StartAPP()

	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	} else {
		status = models.STATUS_PLAYING
	}

	t.updateStatus(status, online)
	return err
}

func (t *Edge) StopAPP() error {

	status, _ := t.GetCacheStatus()
	if status != models.STATUS_PLAYING {
		return errDef.ErrNotPlaying
	}

	status = models.STATUS_RX_STOP_APP
	t.updateStatus(status, nil)

	err := t.eHttp.StopAPP()

	var online *bool
	if err != nil {
		if err == errDef.ErrEdgeLost {
			tmp := false
			online = &tmp
		} else {
			status = models.STATUS_FAIL
		}
	} else {
		status = models.STATUS_RESERVE_XR_CONNECT
	}

	t.updateStatus(status, online)

	return err
}

func (t *Edge) updateStatus(status models.EdgeStatus, online *bool) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.Status = status

	if online != nil {
		t.Online = *online
	}
}

func (t *Edge) setOnline(online bool) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.Online = online

}
