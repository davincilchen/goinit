package errordef

import "errors"

var (
	ErrNoResource = errors.New("no resource")
	ErrEdgeLost   = errors.New("edge lost")

	ErrRepeatedLogin   = errors.New("repeated login")
	ErrRepeatedReserve = errors.New("repeated reserve")
	ErrDevNoReserve    = errors.New("device no reserve")
	ErrInOldProcess    = errors.New("still in old process")

	ErrStartAppTimeout = errors.New("start app timeout")
	ErrInvalidStramVR  = errors.New("invalid steam VR")
	ErrCloudXRUnconect = errors.New("cloudXR unconnect")

	ErrNotPlaying     = errors.New("not playing")
	ErrAlreadyPlaying = errors.New("already playing")
	ErrAlreadyFree    = errors.New("already free")
	ErrProcessing     = errors.New("processing")

	ErrUrlParamError = errors.New("url param error")
)
