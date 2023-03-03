package server

import (
	"xr-central/pkg/delivery"

	"github.com/gin-gonic/gin"

	edgedlv "xr-central/pkg/app/edge/delivery"
	loginDlv "xr-central/pkg/app/login/delivery"
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
	devSession.POST("/apps/:id/reserve", edgedlv.NewOrder)
	devSession.DELETE("/reserve", edgedlv.ReleaseOrder)
	devSession.GET("/resume", edgedlv.DeviceResume)
	devSession.POST("/start_app", edgedlv.StartApp)
	devSession.POST("/stop_app", edgedlv.StopApp)
	devSession.POST("/status", edgedlv.EdgeStatus)

	// .. //
	apps := router.Group("/apps")
	apps.GET("/", delivery.AppList)
	return router
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
