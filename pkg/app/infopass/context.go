package infopass

import (
	"context"
)

// =========================================== //
//最大集合
type Context interface {
	context.Context
	Cache
}

//子集合
type ContextD interface {
	context.Context
	DBErrCache
}

//子集合
type ContextH interface {
	context.Context
	HttpErrCache
}

//當作input參數傳遞
type ContextWithCache interface {
	context.Context
	InfoCache
}

// =========================================== //

type InfoContext struct {
	context.Context
	DBErrPass
	HttpErrPass
}

func NewContext(context ContextWithCache) *InfoContext {

	return &InfoContext{
		Context:     context,
		DBErrPass:   *NewDBErrPass(context),
		HttpErrPass: *NewHttpErrPass(context),
	}

}
