package controllers

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
