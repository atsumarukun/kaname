package routes

import (
	"backend/internal/app/api/computers"

	"github.com/gin-gonic/gin"
)

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

	computers.AddComputerRoutes(v1)
}
