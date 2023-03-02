package errordef

import "errors"

var (
	ErrRepeatedLogin = errors.New("repeated login")
)
