package computers

import (
	"backend/internal/app/api/computers/infrastructures/persistences"
	"backend/internal/app/api/computers/infrastructures/temporaries"
	"backend/internal/app/api/computers/interfaces/handlers"
	"backend/internal/app/api/computers/usecases"
	"backend/internal/app/api/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/mdlayher/wol"
)

func AddComputerRoutes(rg *gin.RouterGroup) {
	wolClient, err := wol.NewClient()
	if err != nil {
		panic(err.Error())
	}

	computerTemporary := temporaries.NewComputerTemporary(wolClient)
	computerPersistence := persistences.NewComputerPersistence()
	computerUsecase := usecases.NewComputerUsecase(computerTemporary, computerPersistence, database.Get())
	computerHandler := handlers.NewComputerHandler(computerUsecase)

	r := rg.Group("computers")
	r.GET("/", computerHandler.SearchComputers)
	r.POST("/", computerHandler.CreateComputer)
	r.GET("/:id", computerHandler.GetComputer)
	r.PUT("/:id", computerHandler.UpdateComputer)
	r.DELETE("/:id", computerHandler.DeleteComputer)
	r.PUT("/:id/wake", computerHandler.WakeComputer)
}
