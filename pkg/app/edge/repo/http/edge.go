package http

// import "xr-central/pkg/models"
import (
	"fmt"
	httph "xr-central/pkg/httphelper"
)

type Edge struct {
	URL string
}

func (t *Edge) SetURL(url string) {
	t.URL = url
}

func (t *Edge) Reserve() error {
	url := fmt.Sprintf("%s//reserve", t.URL)
	_, err := httph.Post(url)
	return err
}

func (t *Edge) Release() error {
	url := fmt.Sprintf("%s//reserve", t.URL)
	_, err := httph.Delete(url)
	return err
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
	url := fmt.Sprintf("%s//%d//start_app", t.URL, appID)
	_, err := httph.Post(url)
	return err
}

func (t *Edge) StopAPP() error {
	url := fmt.Sprintf("%s//stop_app", t.URL)
	_, err := httph.Post(url)

	return err
	//return errDef.ErrEdgeLost
	//return nil
}
