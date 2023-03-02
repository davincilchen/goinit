package delivery

// func getBodyForUserLogin(ctx *gin.Context) (*UserLoginParams, error) {

// 	body, err := GetRawData(ctx) //GetRawDataAndCacheInGin(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var info UserLoginParams
// 	err = json.Unmarshal(body, &info)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &info, nil
// }

// func AuthWhenPlayerLogin(ctx *gin.Context) {
// 	param, err := getBodyForUserLogin(ctx)
// 	response := ResBody{}
// 	if err != nil {
// 		response.ResCode = RES_ERROR_BAD_REQUEST
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	userUCase := usecase.User{}
// 	qRet, err := userUCase.Login(*param.Account, *param.Password)
// 	if err != nil {
// 		response.ResCode = RES_INVALID_USER_PASSWORD
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	data := LoginResponse{
// 		ID:    qRet.ID,
// 		Name:  qRet.Name,
// 		Token: qRet.Token,
// 	}
// 	response.ResCode = RES_OK
// 	response.Data = data
// 	ctx.JSON(http.StatusOK, response)
// }

// .. //
