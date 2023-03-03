package server

import (
	"fmt"
	"xr-central/pkg/app/infopass"
	dlv "xr-central/pkg/delivery"

	"github.com/gin-gonic/gin"
)

// func ResponseRequestError(ctx *gin.Context, err error) {
// 	CacheErrorInGin(ctx, err)
// 	response := MakeResponseWithError(RES_ERROR_BAD_REQUEST, "", err.Error())

// 	CacheResponseInGin(ctx, response)
// 	code := http.StatusBadRequest
// 	CacheHttpStatusCodeInGin(ctx, code)
// 	ctx.JSON(code, response)
// 	ctx.Abort()

// }

// func ResponseAuthError(ctx *gin.Context, err *AuthError) {
// 	CacheErrorInGin(ctx, err)
// 	response := MakeResponseWithError(RES_ERROR_BAD_REQUEST,
// 		"", err.DescForResponse)
// 	if err.Code == AUTH_ERROR_AUTH_SESSION_TOKEN || err.Code == AUTH_ERROR_AUTH_FAILED {
// 		response.ResCode = RES_USER_INVALID_TOKEN
// 	} else if err.Code == AUTH_ERROR_AUTH_PLAYER_TOKEN {
// 		response.ResCode = RES_AUTHENTICATE_TOKEN_FAILED
// 	}

// 	CacheResponseInGin(ctx, response)
// 	code := http.StatusUnauthorized
// 	CacheHttpStatusCodeInGin(ctx, code)
// 	ctx.JSON(code, response)
// 	ctx.Abort()

// }

// ===================== //

// func AuthWhenDevelopLogin(ctx *gin.Context) {
// 	login, err := getBodyDevelopLogin(ctx)
// 	if err != nil {
// 		ResponseRequestError(ctx, err)
// 		return
// 	}

// 	session, errAuth := manager.AuthDevelop(login)
// 	if errAuth != nil {
// 		ResponseAuthError(ctx, errAuth)
// 		return
// 	}

// 	ctx.Set(GinKeySessionToken, session.GetUUID())
// 	cachePlayerSessionInGin(ctx, session)
// }

// func AuthWhenPlayerLogin(ctx *gin.Context) {
// 	login, err := getBodyLogin(ctx)
// 	if err != nil {
// 		ResponseRequestError(ctx, err)
// 		return
// 	}
// 	ctx.Set(GinKeyLoginInfo, login)

// 	supplier, errAuth := AuthSupplier(&login.LoginSupplier)

// 	if errAuth != nil {
// 		ResponseAuthError(ctx, errAuth)
// 		return
// 	}

// 	CacheSupplierInGin(ctx, supplier)

// }

func getSessionToken(ctx *gin.Context) (string, error) {

	key := "authorization"
	// _, ok := ctx.Request.Header[key] //authorization 取不到//"X-Session-Token"可以
	// if !ok {
	// 	return "", fmt.Errorf("getSessionToken failed")
	// }

	authorization := ctx.GetHeader(key)
	if authorization == "" {
		return "", fmt.Errorf("getSessionToken failed")
	}

	// s := strings.Split(authorization, "Bearer ")
	// if len(s) < 2 {
	// 	return "", fmt.Errorf("Split Bearer failed form [%s]", authorization)
	// }
	// token := s[1]

	//return token, nil

	return authorization, nil
}

func AuthDevSession(ctx *gin.Context) {
	sessionToken, err := getSessionToken(ctx)
	if err != nil {
		dlv.RespUnauthorized(ctx, err)
		return
	}
	infopass.CacheSessionToken(ctx, sessionToken)
}

// func auth(ctx *gin.Context, allowOffline bool) {
// 	sessionToken, err := getSessionTokenInAuth(ctx)
// 	if err != nil {
// 		ResponseRequestError(ctx, err)
// 		return
// 	}
// 	CacheSessionTokenInGin(ctx, sessionToken)

// 	loginPlayer, errAuth := AuthLoginToken(sessionToken, allowOffline)
// 	if errAuth != nil {
// 		ResponseAuthError(ctx, errAuth)
// 		return
// 	}
// 	CachePlayerSessionInGin(ctx, loginPlayer)
// }

// func authBlockPlayer() {
// }

// func authDuplicatePlayer() {

// }
