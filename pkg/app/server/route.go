package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	appDlv "goinit/pkg/app/app/delivery"
	devDlv "goinit/pkg/app/device/delivery"
	edgeDlv "goinit/pkg/app/edge/delivery"
	loginDlv "goinit/pkg/app/login/delivery"
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

	//
	router.StaticFS("/static", http.Dir("web/static"))

	// .. //
	dev := router.Group("/devices")
	dev.GET("/", devDlv.DeviceList)
	dev.POST("/login", loginDlv.DevLogin)

	devSession := router.Group("/devices")
	devSession.Use(AuthDevSession)

	devSession.POST("/logout", loginDlv.DevLogout)
	devSession.POST("/apps/:id/reserve", edgeDlv.NewReserve)
	devSession.DELETE("/reserve", edgeDlv.ReleaseReserve)
	devSession.GET("/resume", edgeDlv.DeviceResume)
	devSession.POST("/start_app", edgeDlv.StartApp)
	devSession.POST("/stop_app", edgeDlv.StopApp)
	devSession.POST("/status", edgeDlv.EdgeStatus)

	// .. //
	apps := router.Group("/apps")
	apps.GET("/", appDlv.AppList)

	// .. //
	edges := router.Group("/edges")
	edges.GET("/", edgeDlv.EdgeList)
	edges.POST("/reg", edgeDlv.EdgeReg)
	edges.GET("/:id/apps", edgeDlv.EdgeAppList)

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
