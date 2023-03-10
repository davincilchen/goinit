package http

// import "xr-central/pkg/models"
import (
	"fmt"
	"net/http"
	"xr-central/pkg/app/ctxcache"
	errDef "xr-central/pkg/app/errordef"
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

func (t *Edge) Reserve(ctx ctxcache.Context, appID int) error {
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

func (t *Edge) Release(ctx ctxcache.Context) error {
	url := fmt.Sprintf("http://%s//reserve", t.URL)
	_, err := httph.Delete(url)
	return err
}

func (t *Edge) Resume(ctx ctxcache.Context) error {

	return nil
}

func (t *Edge) GetStatus(ctx ctxcache.Context) error {

	return nil
}

func (t *Edge) Status(ctx ctxcache.Context) error {

	return nil
}

func (t *Edge) StartAPP(ctx ctxcache.Context, appID int) error {
	url := fmt.Sprintf("http://%s//%d//start_app", t.URL, appID)
	_, err := httph.Post(url)
	return err
}

func (t *Edge) StopAPP(ctx ctxcache.Context) error {
	url := fmt.Sprintf("http://%s//stop_app", t.URL)
	_, err := httph.Post(url)

	return err
	//return errDef.ErrEdgeLost
	//return nil
}
