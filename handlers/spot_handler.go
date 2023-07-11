package handlers

import (
	"net/http"
	"strconv"

	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/services"
	"github.com/gin-gonic/gin"
)

type SpotHandler struct {
	spotService services.SpotService
}

func NewSpotHandler(spotService services.SpotService) *SpotHandler {
	return &SpotHandler{spotService: spotService}
}

func (sh *SpotHandler) CreateSpot(c *gin.Context) {
	var spot models.Spot
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spot data"})
		return
	}

	if err := sh.spotService.CreateSpot(&spot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create spot"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spot created successfully"})
}

func (sh *SpotHandler) GetSpot(c *gin.Context) {
	spotIDStr := c.Param("id")
	spotID, err := strconv.ParseUint(spotIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spot ID"})
		return
	}

	spot, err := sh.spotService.GetSpot(uint(spotID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get spot"})
		return
	}

	c.JSON(http.StatusOK, spot)
}

func (sh *SpotHandler) GetAllSpots(c *gin.Context) {
	spots, err := sh.spotService.GetAllSpots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spots"})
		return
	}
	c.JSON(http.StatusOK, spots)
}
