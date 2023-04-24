package delivery

import (
	"net/http"

	"xr-central/pkg/app/ctxcache"

	"github.com/gin-gonic/gin"

	appUCase "xr-central/pkg/app/app/usecase"
	dlv "xr-central/pkg/delivery"
)

type AppListResp struct {
	Total int       `json:"total_num"`
	Apps  []AppResp `json:"app"`
}

type AppResp struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Genre          string `json:"genre"`
	Des            string `json:"description"`
	ImgURL         string `json:"img_url"`
	CentralPathImg bool   `json:"central_path_img"`
}

func AppList(ctx *gin.Context) { //TODO:
	appHandle := appUCase.AppHandle{}
	apps, err := appHandle.GetApps(true)
	if err != nil {
		ctxcache.CacheDBError(ctx, err)
		dlv.RespError(ctx, err, nil)
		return
	}

	data := AppListResp{
		Total: len(apps),
	}

	for _, v := range apps {
		a := AppResp{
			ID:             int(v.ID),
			Title:          v.AppTitle,
			Genre:          v.AppGenre.Type,
			Des:            v.AppBrief,
			ImgURL:         v.ImageURL,
			CentralPathImg: v.CentralImage,
		}
		data.Apps = append(data.Apps, a)
	}

	response := dlv.ResBody{}
	response.ResCode = dlv.RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
