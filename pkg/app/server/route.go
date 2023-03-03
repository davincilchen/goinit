package server

import (
	loginDlv "xr-central/pkg/app/login/delivery"
	"xr-central/pkg/delivery"

	"github.com/gin-gonic/gin"
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
	devSession.POST("/apps/:id/reserve", delivery.NewOrder)
	devSession.DELETE("/reserve", delivery.ReleaseOrder)
	devSession.GET("/resume", delivery.DeviceResume)
	devSession.POST("/start_app", delivery.StartApp)
	devSession.POST("/stop_app", delivery.StopApp)
	devSession.POST("/status", delivery.EdgeStatus)

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
