package http

// import "xr-central/pkg/models"
import (
	errDef "xr-central/pkg/app/errordef"
)

type Edge struct {
	URL string
}

func (t *Edge) SetURL(url string) {
	t.URL = url
}

func (t *Edge) Reserve() error {
	return nil
}

func (t *Edge) Release() error {

	return nil
}

func (t *Edge) Resume() error {

	return nil
}

func (t *Edge) GetStatus() error {

	return nil
}

func (t *Edge) Status() error {

	return nil
}

func (t *Edge) StartAPP(appID int) error {

	return nil
}

func (t *Edge) StopAPP() error {

	return errDef.ErrEdgeLost
	//return nil
}
