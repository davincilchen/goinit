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

	status := models.STATUS_FREE
	err := t.eHttp.Release()

	var online *bool
	if err == errDef.ErrEdgeLost {
		tmp := false
		online = &tmp
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

func (t *Edge) StartAPP() error {
	status := models.STATUS_START_APP // or

	err := t.eHttp.StartAPP()

	var online *bool
	if err == errDef.ErrEdgeLost {
		tmp := false
		online = &tmp
	}

	t.updateStatus(status, online)
	return err
}

func (t *Edge) StopAPP() error {

	status := models.STATUS_RESERVE_XR_CONNECT // or
	err := t.eHttp.StopAPP()

	var online *bool
	if err == errDef.ErrEdgeLost {
		tmp := false
		online = &tmp
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
