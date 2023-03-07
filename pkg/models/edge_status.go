package models

type EdgeStatus int

const (
	STATUS_FREE         EdgeStatus = 0
	STATUS_RESERVE_INIT EdgeStatus = 110
	//STATUS_RESERVE_PROCESSS       EdgeStatus = 120
	STATUS_RESERVE_XR_NOT_CONNECT EdgeStatus = 130
	STATUS_RESERVE_XR_CONNECT     EdgeStatus = 140
	STATUS_RX_START_APP           EdgeStatus = 150
	//STATUS_APP_RUNNING            EdgeStatus = 160
	STATUS_PLAYING     EdgeStatus = 170
	STATUS_RX_STOP_APP EdgeStatus = 180
	STATUS_RX_RELEASE  EdgeStatus = 190

	STATUS_FAIL EdgeStatus = 999
	//suspend?
)
