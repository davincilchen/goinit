package server

import (
	"central/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var show = "show nothing"

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(Logger, gin.Recovery())

	router.GET("/hello", hello)
	router.POST("/login", controllers.Login)

	// .. //
	user := router.Group("/user") //本人
	user.POST("/deposit", controllers.Deposit)
	user.POST("/withdrawal", controllers.Withdraw)

	users := router.Group("/users")
	users.POST("/:id/deposit", controllers.DepositTo)
	users.POST("/:id/withdrawal", controllers.WithdrawFrom)

	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{ // response json
		"message": "Hello" + ",  " + show,
	})
}
