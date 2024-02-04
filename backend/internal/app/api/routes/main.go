package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run(addr ...string) {
	getApiRoutes()
	router.Run(addr...)
}

func getApiRoutes() {
	api := router.Group("api")
	addV1Routes(api)
}

func addV1Routes(rg *gin.RouterGroup) {
	v1 := rg.Group("v1")

	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}