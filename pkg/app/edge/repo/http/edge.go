package http

// import "xr-central/pkg/models"
import (
	"fmt"
	"net/http"
	errDef "xr-central/pkg/app/errordef"
	"xr-central/pkg/app/infopass"
	httph "xr-central/pkg/httphelper"
)

func NewEdge(URL string, errCache infopass.HttpErrCache) *Edge {
	return &Edge{
		URL:      URL,
		errCache: errCache,
	}

}

type Edge struct {
	URL      string
	errCache infopass.HttpErrCache
}

func (t *Edge) SetURL(url string) {
	t.URL = url
}

func (t *Edge) Reserve(appID int) error {
	url := fmt.Sprintf("http://%s//reserve//app//%d", t.URL, appID)
	resp, err := httph.Post(url)
	if err != nil {
		fmt.Println(err)
		t.errCache.CacheHttpError(err)
		return errDef.ErrEdgeLost
	}
	fmt.Println(resp)
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("resp.StatusCode = %d", resp.StatusCode)
		t.errCache.CacheHttpError(err)
		return err
	}

	return nil
}

func (t *Edge) Release() error {
	url := fmt.Sprintf("http://%s//reserve", t.URL)
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
	url := fmt.Sprintf("http://%s//%d//start_app", t.URL, appID)
	_, err := httph.Post(url)
	return err
}

func (t *Edge) StopAPP() error {
	url := fmt.Sprintf("http://%s//stop_app", t.URL)
	_, err := httph.Post(url)

	return err
	//return errDef.ErrEdgeLost
	//return nil
}
