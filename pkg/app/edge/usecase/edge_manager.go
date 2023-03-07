package usecase

import "xr-central/pkg/models"

type EdgeManager struct {
}

var manager *EdgeManager

func newEdgeManager() *EdgeManager {
	d := &EdgeManager{}

	return d
}

func GetEdgeManager() *EdgeManager {
	if manager == nil {
		manager = newEdgeManager()
	}
	return manager
}

func (t *EdgeManager) Reserve(appID int) (*Edge, error) {

	edge := &Edge{}
	//don't need lock it's new
	edge.Status = models.STATUS_RESERVE_INIT //lock
	//edge.Status = models.STATUS_RESERVE_PROCESSS       //lock
	edge.Status = models.STATUS_RESERVE_XR_NOT_CONNECT //lock
	return nil, nil
}
