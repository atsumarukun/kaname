package computers

import (
	"backend/internal/app/api/computers/infrastructure/persistences"
	"backend/internal/app/api/computers/interface/handlers"
	"backend/internal/app/api/computers/usecases"
	"backend/internal/app/api/database"

	"github.com/gin-gonic/gin"
)

func AddComputerRoutes(rg *gin.RouterGroup) {
	computerPersistence := persistences.NewComputerPersistence()
	computerUsecase := usecases.NewComputerUsecase(computerPersistence, database.Get())
	computerHandler := handlers.NewComputerHandler(computerUsecase)

	r := rg.Group("computers")
	r.GET("/", computerHandler.SearchComputers)
	r.POST("/", computerHandler.CreateComputer)
	r.GET("/:id", computerHandler.GetComputer)
	r.PUT("/:id", computerHandler.UpdateComputer)
	r.DELETE("/:id", computerHandler.DeleteComputer)
	r.PUT("/:id/wake", computerHandler.WakeComputer)
}
