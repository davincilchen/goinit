package usecase

import (
	"fmt"
	repo "xr-central/pkg/app/edge/repo/mysql"
	"xr-central/pkg/models"
)

type EdgeManager struct {
	//TODO: lock
	edges []*Edge
}

var manager *EdgeManager

func newEdgeManager() *EdgeManager {
	d := &EdgeManager{}

	return d
}

func GetEdgeManager() *EdgeManager {
	if manager == nil {
		manager = newEdgeManager()
		manager.edges = make([]*Edge, 0)
		e := repo.Edge{}
		es, err := e.LoadEdges()
		if err != nil {
			fmt.Printf("LoadEdges error %s\n", err.Error())
		} else {
			fmt.Printf("LoadEdges count %d\n", len(es))
			for i, v := range es {
				fmt.Printf("%d %#v\n", i, v)
				manager.edges = append(manager.edges, NewEdge(v))
			}
		}

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
