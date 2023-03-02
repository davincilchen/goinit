package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResCode int

const (
	RES_OK ResCode = 0

	RES_NO_RESOURCE       ResCode = 100
	RES_EDGE_LOST         ResCode = 101
	RES_INVALID_STEAM_VR  ResCode = 102
	RES_CLOUDXR_UNCONNECT ResCode = 103

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

func RespBadRequest(ctx *gin.Context) {

	response := ResBody{}
	response.ResCode = RES_ERROR_BAD_REQUEST
	ctx.JSON(http.StatusBadRequest, response)

}

func RespInvalidPassword(ctx *gin.Context) {

	response := ResBody{}
	response.ResCode = RES_INVALID_USER_PASSWORD
	ctx.JSON(http.StatusBadRequest, response)

}

func RespUnknowError(ctx *gin.Context, err error) {

	response := ResBody{}
	response.ResCode = RES_ERROR_UNKNOWN
	ctx.JSON(http.StatusInternalServerError, response)

}
