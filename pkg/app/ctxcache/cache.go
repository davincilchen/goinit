package ctxcache

import (
	"context"
	"fmt"
)

type any = interface{}

const (
	GinKeyError          = "Error"
	GinKeyAdvError       = "AdvError"
	GinKeyDBError        = "DBError"
	GinKeyHttpError      = "HttpError"
	GinKeyInfo           = "Info"
	GinKeyLoginInfo      = "LoginInfo"
	GinKeyPlayerSession  = "PlayerSession"
	GinKeySessionToken   = "SessionToken"
	GinKeyRequestBody    = "RequestBody"
	GinKeyHttpStatusCode = "HttpStatusCode"
	GinKeyResponse       = "Response"
	GinKeyHandleContext  = "HandleContext"
	//GinKeyDevice         = "Device"
)

// =========================================== //
type InfopassContent interface {
	context.Context
}

// =========================================== //

type InfoCache interface {
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}

type DBErrCache interface {
	CacheDBError(error)
	GetDBError() error
}

type HttpErrCache interface {
	ResetHttpError()
	CacheHttpError(error)
	GetHttpError() error
}

type Cache interface {
	DBErrCache
	HttpErrCache
}

// =========================================== //

func NewDBErrPass(cache InfoCache) *DBErrPass {
	return &DBErrPass{
		cache: cache,
	}
}

type DBErrPass struct {
	cache InfoCache
}

func (t *DBErrPass) CacheDBError(err error) {
	CacheDBError(t.cache, err)
}

func (t *DBErrPass) GetDBError() error {
	return GetDBError(t.cache)
}

// ========================== //
func NewHttpErrPass(cache InfoCache) *HttpErrPass {
	return &HttpErrPass{
		cache: cache,
	}
}

type HttpErrPass struct {
	cache InfoCache
}

func (t *HttpErrPass) ResetHttpError() {
	ResetHttpError(t.cache)
}

func (t *HttpErrPass) CacheHttpError(err error) {
	CacheHttpError(t.cache, err)
}

func (t *HttpErrPass) GetHttpError() error {
	return GetHttpError(t.cache)
}

// ================================ //

func CacheError(ctx InfoCache, err error) {
	if err == nil {
		return
	}
	ctx.Set(GinKeyError, err)
}

func GetError(ctx InfoCache) error {
	err, exist := ctx.Get(GinKeyError)
	if !exist {
		return nil
	}
	if err != nil {
		e, ok := err.(error)
		if ok {
			return e
		}
	}
	return nil
}

func CacheAdvError(ctx InfoCache, err error) {
	if err == nil {
		return
	}
	ctx.Set(GinKeyAdvError, err)
}

func GetAdvError(ctx InfoCache) error {
	err, exist := ctx.Get(GinKeyAdvError)
	if !exist {
		return nil
	}
	if err != nil {
		e, ok := err.(error)
		if ok {
			return e
		}
	}
	return nil
}

func CacheDBError(ctx InfoCache, err error) {
	if err == nil {
		return
	}
	ctx.Set(GinKeyDBError, err)
}

func GetDBError(ctx InfoCache) error {
	err, exist := ctx.Get(GinKeyDBError)
	if !exist {
		return nil
	}
	if err != nil {
		e, ok := err.(error)
		if ok {
			return e
		}
	}
	return nil
}

func ResetHttpError(ctx InfoCache) {
	ctx.Set(GinKeyHttpError, nil)
}

func CacheHttpError(ctx InfoCache, err error) {
	if err == nil {
		return
	}
	ctx.Set(GinKeyHttpError, err)
}

func GetHttpError(ctx InfoCache) error {
	err, exist := ctx.Get(GinKeyHttpError)
	if !exist {
		return nil
	}
	if err != nil {
		e, ok := err.(error)
		if ok {
			return e
		}
	}
	return nil
}

func CacheSessionToken(ctx InfoCache, sessionToken string) {
	ctx.Set(GinKeySessionToken, sessionToken)
}

func GetSessionToken(ctx InfoCache) *string {
	ret, exist := ctx.Get(GinKeySessionToken)
	if !exist {
		return nil
	}
	if ret != nil {
		e, ok := ret.(string)
		if ok {
			return &e
		}
	}
	return nil
}

func CacheRequestBodyInGin(ctx InfoCache, requestBody *string) {
	ctx.Set(GinKeyRequestBody, requestBody)
}

func GetRequestBodyInGin(ctx InfoCache) (*string, error) {

	info, ok := ctx.Get(GinKeyRequestBody)
	if !ok {
		err := fmt.Errorf("GinKeyRequestBody is not find")
		return nil, err
	}

	r, ok := info.(*string)
	if !ok {
		err := fmt.Errorf("trans GinKeyRequestBody.(*string) failed from cache")
		return nil, err
	}

	return r, nil
}
