package deliverymodel

import (
	"xr-central/pkg/models"

	// "encoding/json"
	// "errors"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strconv"

	// "github.com/gin-gonic/gin"

	// "xr-central/pkg/app/ctxcache"
	//devUCase "xr-central/pkg/app/device/usecase"
	edgeUCase "xr-central/pkg/app/edge/usecase"
)

// errDef "xr-central/pkg/app/errordef"
// dlv "xr-central/pkg/delivery"
// "xr-central/pkg/models"

type EdgeInfo struct {
	ID     uint                `json:"id"`
	IP     string              `json:"ip"`
	Port   int                 `json:"port"`
	Status models.EdgeStatus   `json:"status"`
	Online bool                `json:"online"`
	ActRet edgeUCase.ActionRet `json:"last_act_ret"`

	Device *DeviceInfo `json:"device"`
	AppID  *uint       `json:"app_id"`
}
