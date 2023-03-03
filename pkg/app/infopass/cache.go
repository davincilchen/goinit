package infopass

import (
	"fmt"
)

type any = interface{}

type InfoCache interface {
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}

const (
	GinKeyError          = "Error"
	GinKeyDBError        = "DBError"
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
