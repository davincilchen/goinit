package errordef

import "errors"

var (
	ErrNoResource      = errors.New("no resource")
	ErrEdgeLost        = errors.New("edge lost")
	ErrStartAppTimeout = errors.New("start app timeout")
	ErrInvalidStramVR  = errors.New("invalid steam VR")
	ErrCloudXRUnconect = errors.New("cloudXR unconnect")
	ErrNotPlaying      = errors.New("not playing")
	ErrAlreadyFree     = errors.New("already free")

	ErrRepeatedLogin   = errors.New("repeated login")
	ErrRepeatedReserve = errors.New("repeated reserve")
	ErrDevNoReserve    = errors.New("device no reserve")

	ErrUrlParamError = errors.New("url param error")
)
