package handlers

import (
	"net/http"

	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/services"
	"github.com/gin-gonic/gin"
)

type GateHandler struct {
	gateService services.GateService
}

func NewGateHandler(gateService services.GateService) *GateHandler {
	return &GateHandler{gateService: gateService}
}

func (gh *GateHandler) CreateGateEntry(c *gin.Context) {
	var gateEntry models.GateEntry
	if err := c.ShouldBindJSON(&gateEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gate entry data"})
		return
	}

	if err := gh.gateService.CreateGateEntry(&gateEntry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gate entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gate entry created successfully"})
}

func (gh *GateHandler) GetTotalGateCount(c *gin.Context) {
	count, err := gh.gateService.GetTotalGateCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total gate count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (gh *GateHandler) GetGateCountByGate(c *gin.Context) {

	count, err := gh.gateService.GetGateCountByGate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gate count"})
		return
	}

	c.JSON(http.StatusOK, count)
}

func (gh *GateHandler) GetGates(c *gin.Context) {
	gates, err := gh.gateService.GetGates()
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gates"})
		return
	}
	c.JSON(http.StatusOK, gates)
}

