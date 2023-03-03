package usecase

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

	return nil, nil
}
