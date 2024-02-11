package handlers

import (
	"backend/internal/app/api/computers/interface/requests"
	"backend/internal/app/api/computers/interface/responses"
	"backend/internal/app/api/computers/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComputerHandler interface {
	CreateComputer(*gin.Context)
}

type computerHandler struct {
	computerUsecase usecases.ComputerUsecase
}

func NewComputerHandler(cu usecases.ComputerUsecase) ComputerHandler {
	return &computerHandler{
		computerUsecase: cu,
	}
}

func (ch computerHandler) CreateComputer(c *gin.Context) {
	var input requests.CreateComputerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ch.computerUsecase.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntity(entity)})
}
