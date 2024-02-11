package handlers

import (
	"backend/internal/app/api/computers/interface/requests"
	"backend/internal/app/api/computers/interface/responses"
	"backend/internal/app/api/computers/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComputerHandler interface {
	CreateComputer(*gin.Context)
	UpdateComputer(*gin.Context)
	DeleteComputer(*gin.Context)
	GetComputer(*gin.Context)
	SearchComputers(*gin.Context)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntity(entity)})
}

func (ch computerHandler) UpdateComputer(c *gin.Context) {
	var input requests.UpdateComputerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ch.computerUsecase.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntity(entity)})
}

func (ch computerHandler) DeleteComputer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ch.computerUsecase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntity(entity)})
}

func (ch computerHandler) GetComputer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ch.computerUsecase.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntity(entity)})
}

func (ch computerHandler) SearchComputers(c *gin.Context) {
	var query requests.SearchComputersQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entities, err := ch.computerUsecase.Search(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses.FromEntities(entities)})
}
