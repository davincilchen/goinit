package server

import (
	login "xr-central/pkg/app/login/delivery"
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
	router.POST("/login", login.Login)
	router.POST("/logout", login.Logout)

	// .. //
	edges := router.Group("/devices")
	edges.POST("/login", login.DevLogin)
	edges.POST("/logout", login.Logout)

	edgesSession := router.Group("/devices")
	edgesSession.Use(AuthDevSession)

	edgesSession.POST("/apps/:id/reserve", delivery.NewOrder)
	edgesSession.DELETE("/reserve", delivery.ReleaseOrder)
	edgesSession.GET("/resume", delivery.DeviceResume)
	edgesSession.POST("/start_app", delivery.StartApp)
	edgesSession.POST("/stop_app", delivery.StopApp)
	edgesSession.POST("/status", delivery.EdgeStatus)

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
