package computers

import (
	"backend/internal/app/api/computers/infrastructure/persistences"
	"backend/internal/app/api/computers/interface/handlers"
	"backend/internal/app/api/computers/usecases"

	"github.com/gin-gonic/gin"
)

func AddComputerRoutes(rg *gin.RouterGroup) {
	computerPersistence := persistences.NewComputerPersistence()
	computerUsecase := usecases.NewComputerUsecase(computerPersistence)
	computerHandler := handlers.NewComputerHandler(computerUsecase)

	r := rg.Group("computers")
	r.POST("/", computerHandler.CreateComputer)
}