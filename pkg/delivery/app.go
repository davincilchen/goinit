package delivery

import (

	//"xr-central/pkg/app/users/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

// func getBodyForLogin(ctx *gin.Context) (*LoginParams, error) {

// 	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var info LoginParams
// 	err = json.Unmarshal(body, &info)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &info, nil
// }

type AppListResp struct {
	Total int `json:"total_num"`
}

func AppList(ctx *gin.Context) { //TODO:
	data := AppListResp{}

	response := ResBody{}
	response.ResCode = RES_OK
	response.Data = data

	ctx.JSON(http.StatusOK, response)
}
