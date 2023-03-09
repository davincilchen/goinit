package infopass

import (
	"fmt"
)

type any = interface{}

type InfoCache interface {
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}

type DBErrCache interface {
	CacheDBError(error)
	GetDBError() error
}

type HttpErrCache interface {
	CacheHttpError(error)
	GetHttpError() error
}

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

func NewDBErrorProc(porc InfoCache) *DBErrorProc {
	return &DBErrorProc{
		porc: porc,
	}
}

type DBErrorProc struct {
	porc InfoCache
}

func (t *DBErrorProc) CacheDBError(err error) {
	CacheDBError(t.porc, err)
}

func (t *DBErrorProc) GetDBError() error {
	return GetDBError(t.porc)
}

// ============== //
func NewHttpErrorProc(porc InfoCache) *HttpErrorProc {
	return &HttpErrorProc{
		porc: porc,
	}
}

type HttpErrorProc struct {
	porc InfoCache
}

func (t *HttpErrorProc) CacheHttpError(err error) {
	CacheHttpError(t.porc, err)
}

func (t *HttpErrorProc) GetHttpError() error {
	return GetHttpError(t.porc)
}

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
