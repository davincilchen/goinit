package delivery

// import (
// 	"xr-central/pkg/app/login/delivery"
// 	userUCase "xr-central/pkg/app/users/usecase"

// 	"github.com/gin-gonic/gin"
// )

// func DevLogin(ctx *gin.Context) {
// 	req := &delivery.DevLoginParams{}
// 	err := GetBodyFromRawData(ctx, req)
// 	if err != nil {
// 		RespBadRequest(ctx)
// 		return
// 	}

// 	handle := delivery.NewLoginController(ctx, req.UserLoginParams, DevLoginSucess)
// 	handle.Do()

// }

// func DevLoginSucess(user *userUCase.LoginUser) error {

// 	return nil
// }
