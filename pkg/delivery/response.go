package delivery

import (
	"net/http"

	"xr-central/pkg/app/errordef"
	"xr-central/pkg/app/infopass"

	"github.com/gin-gonic/gin"
)

type ResCode int

const (
	RES_OK ResCode = 0

	RES_NO_RESOURCE       ResCode = 100
	RES_EDGE_LOST         ResCode = 101
	RES_INVALID_STEAM_VR  ResCode = 102
	RES_CLOUDXR_UNCONNECT ResCode = 103
	RES_REPEATED_LOGIN    ResCode = 104

	RES_ERROR_UNKNOWN         ResCode = 200
	RES_ERROR_BAD_REQUEST     ResCode = 201
	RES_INVALID_USER_TOKEN    ResCode = 202
	RES_INVALID_USER_PASSWORD ResCode = 203
)

type ResError struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

type ResBody struct {
	ResCode ResCode     `json:"resp_code"`
	Error   *ResError   `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespBadRequest(ctx *gin.Context, err error) {

	response := FillErrorBody(ctx, err)
	response.ResCode = RES_ERROR_BAD_REQUEST
	infopass.CacheError(ctx, err)
	ctx.JSON(http.StatusBadRequest, response)

}

func RespInvalidToken(ctx *gin.Context, err error) {

	response := FillErrorBody(ctx, err)
	response.ResCode = RES_INVALID_USER_TOKEN
	infopass.CacheError(ctx, err)
	ctx.JSON(http.StatusBadRequest, response)
	ctx.Abort()
}

func RespInvalidPassword(ctx *gin.Context) {

	response := FillErrorBody(ctx, nil)
	response.ResCode = RES_INVALID_USER_PASSWORD
	ctx.JSON(http.StatusBadRequest, response)

}

func getStatusCode(err error) (ResCode, int) {

	//logrus.Error(err)
	switch err {
	case errordef.ErrRepeatedLogin:
		return RES_REPEATED_LOGIN, http.StatusBadRequest
	default:
		return RES_ERROR_UNKNOWN, http.StatusInternalServerError
	}
}

func RespUnknowError(ctx *gin.Context, err error) {

	response := FillErrorBody(ctx, err)
	resCode, httpCode := getStatusCode(err)
	response.ResCode = resCode
	infopass.CacheError(ctx, err)

	ctx.JSON(httpCode, response)

}

func FillErrorBody(ctx *gin.Context, err error) *ResBody {
	resp := &ResBody{}
	if err != nil {
		resp.Error = &ResError{
			Desc: err.Error(),
		}
	}
	return resp
}
