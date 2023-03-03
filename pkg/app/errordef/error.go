package errordef

import "errors"

var (
	ErrNoResource      = errors.New("no resource")
	ErrEdgeLost        = errors.New("edge lost")
	ErrInvalidStramVR  = errors.New("invalid steam VR")
	ErrCloudXRUnconect = errors.New("cloudXR unconnect")

	ErrRepeatedLogin   = errors.New("repeated login")
	ErrRepeatedReserve = errors.New("repeated reserve")

	ErrUrlParamError = errors.New("url param error")
)
