package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	appdlv "xr-central/pkg/app/app/delivery"
	edgedlv "xr-central/pkg/app/edge/delivery"
	loginDlv "xr-central/pkg/app/login/delivery"
	"xr-central/pkg/deliveryfake" //just for test. TODO: remove
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(Logger, gin.Recovery())

	router.GET("/exit", exit)
	router.GET("/info", info)
	//router.GET("/edges/status", info)
	//router.GET("/app/usage_satus", info)
	router.POST("/login", loginDlv.Login)
	router.POST("/logout", loginDlv.Logout)

	// .. //
	dev := router.Group("/devices")
	dev.POST("/login", loginDlv.DevLogin)

	devSession := router.Group("/devices")
	devSession.Use(AuthDevSession)

	devSession.POST("/logout", loginDlv.DevLogout)
	devSession.POST("/apps/:id/reserve", edgedlv.NewReserve)
	devSession.DELETE("/reserve", edgedlv.ReleaseReserve)
	devSession.GET("/resume", edgedlv.DeviceResume)
	devSession.POST("/start_app", edgedlv.StartApp)
	devSession.POST("/stop_app", edgedlv.StopApp)
	devSession.POST("/status", edgedlv.EdgeStatus)

	// .. //
	apps := router.Group("/apps")
	apps.GET("/", appdlv.AppList)

	// .. //
	edges := router.Group("/edges")
	edges.GET("/", edgedlv.EdgeList)
	edges.POST("/reg", edgedlv.EdgeReg)

	// .. //fake for test // TODO: remove
	router.POST("/start_app", deliveryfake.FakeStartApp)
	router.POST("/stop_app", deliveryfake.FakeStopApp)

	router.POST("/reserve/app/:id", test)
	return router
}
func test(ctx *gin.Context) {
	fmt.Println(" ============ test ========== ")
}

func info(c *gin.Context) {
	c.JSON(200, gin.H{ // response json
		"version": "0.0.0.1",
	})
}

func exit(c *gin.Context) {
	c.JSON(200, gin.H{ // response json
		"version": "0.0.0.1",
	})
}
