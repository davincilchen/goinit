package errordef

import "errors"

var (
	ErrRepeatedLogin   = errors.New("repeated login")
	ErrRepeatedReserve = errors.New("repeated reserve")
	ErrUrlParamError   = errors.New("url param error")
)
