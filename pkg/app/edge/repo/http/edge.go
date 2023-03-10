package http

// import "xr-central/pkg/models"
import (
	"fmt"
	"net/http"
	errDef "xr-central/pkg/app/errordef"
	"xr-central/pkg/app/infopass"
	httph "xr-central/pkg/httphelper"
)

func NewEdge(URL string) *Edge {
	return &Edge{
		URL: URL,
	}

}

type Edge struct {
	URL string
}

func (t *Edge) SetURL(url string) {
	t.URL = url
}

func (t *Edge) Reserve(ctx infopass.Context, appID int) error {
	url := fmt.Sprintf("http://%s//reserve//app//%d", t.URL, appID)
	resp, err := httph.Post(url)
	if err != nil {
		fmt.Println(err)
		ctx.CacheHttpError(err)
		return errDef.ErrEdgeLost
	}
	fmt.Println(resp)
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("resp.StatusCode = %d", resp.StatusCode)
		ctx.CacheHttpError(err)
		return err
	}

	return nil
}

func (t *Edge) Release(ctx infopass.Context) error {
	url := fmt.Sprintf("http://%s//reserve", t.URL)
	_, err := httph.Delete(url)
	return err
}

func (t *Edge) Resume(ctx infopass.Context) error {

	return nil
}

func (t *Edge) GetStatus(ctx infopass.Context) error {

	return nil
}

func (t *Edge) Status(ctx infopass.Context) error {

	return nil
}

func (t *Edge) StartAPP(ctx infopass.Context, appID int) error {
	url := fmt.Sprintf("http://%s//%d//start_app", t.URL, appID)
	_, err := httph.Post(url)
	return err
}

func (t *Edge) StopAPP(ctx infopass.Context) error {
	url := fmt.Sprintf("http://%s//stop_app", t.URL)
	_, err := httph.Post(url)

	return err
	//return errDef.ErrEdgeLost
	//return nil
}
