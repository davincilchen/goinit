package usecase

import (
	"fmt"
	"sync"
	"xr-central/pkg/app/ctxcache"
	repo "xr-central/pkg/app/edge/repo/mysql"
	"xr-central/pkg/app/errordef"
	"xr-central/pkg/models"
)

type EdgeManager struct {
	//TODO: lock
	edges   []*Edge
	edgeMap map[int]*Edge
	mux     sync.Mutex
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
		manager.edgeMap = make(map[int]*Edge)
		eRepo := repo.Edge{}
		es, err := eRepo.GetEdges()
		if err != nil {
			fmt.Printf("LoadEdges error %s\n", err.Error())
		} else {
			fmt.Println("========= LoadEdges Start =========")
			fmt.Printf("LoadEdges count %d\n", len(es))
			for i, v := range es {
				fmt.Printf("%d %#v\n", i, v)
				tmpEdge := NewEdge(v)
				manager.edgeMap[int(v.ID)] = tmpEdge
				manager.edges = append(manager.edges, tmpEdge)

			}
			fmt.Println("========= LoadEdges Done =========")
		}

	}
	return manager
}

func (t *EdgeManager) Reserve(ctx ctxcache.Context, appID int) (*Edge, error) {
	elist, err := t.FindUnusedEdgesWithAppID(appID)
	if err != nil {
		return nil, err
	}

	var edge *Edge

	for _, v := range elist {
		err := v.Reserve(ctx, appID)
		if err != nil {
			continue
		}
		edge = v
		break
	}

	if edge == nil {
		return nil, errordef.ErrNoResource
	}

	return edge, nil
}

func (t *EdgeManager) FindUnusedEdgesWithAppID(appID int) ([]*Edge, error) {

	eapp, err := t.findEdgeApp(appID)

	if err != nil {
		return nil, err
	}

	edges := make([]*Edge, 0)
	for _, v := range eapp {

		e := t.getEdge(v.EdgeID)
		if e == nil {
			continue
		}

		info := e.GetInfo()
		if !info.Online {
			continue
		}
		if info.Status != models.STATUS_FREE {
			continue
		}
		edges = append(edges, e)
		fmt.Printf("[Reserve List for APP = %d] IP:%s : %d \n", appID, e.info.IP, e.info.Port)
	}

	return edges, nil
}

func (t *EdgeManager) findEdgeApp(appID int) ([]models.EdgeApp, error) {
	eRepo := repo.Edge{}
	edge_app, err := eRepo.FindEdgesWithAppID(appID)

	if err != nil {
		return nil, err
	}
	return edge_app, nil
}

func (t *EdgeManager) getEdge(id int) *Edge {

	t.mux.Lock()
	defer t.mux.Unlock()

	e, ok := t.edgeMap[id]
	if !ok {
		return nil
	}
	return e
}

func (t *EdgeManager) GetEdgeList() []EdgeInfoStatus {

	ret := make([]EdgeInfoStatus, 0)
	for _, v := range t.edges {
		ret = append(ret, v.GetInfo())
	}
	return ret
}
