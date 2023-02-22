package server

import (
	"xr-central/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(Logger, gin.Recovery())

	router.GET("/exit", exit)
	router.GET("/info", info)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	// .. //
	edges := router.Group("/edges")
	edges.POST("/apps/:id/order", controllers.NewOrder)
	edges.DELETE("/order", controllers.ReleaseOrder)
	edges.GET("/resume_app", controllers.ResumeApp)
	edges.POST("/start_app", controllers.StartApp)
	edges.POST("/stop_app", controllers.StopApp)
	edges.GET("/status", controllers.EdgeStatus)
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
