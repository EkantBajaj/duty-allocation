package handlers

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateSpotUserRequest struct {
	SpotID  string    `json:"spot_id" binding:"required"`
	BadgeID string    `json:"badge_id" binding:"required"`
	OutTime time.Time `json:"out_time" binding:"required"`
}
type SpotUserHandler struct {
	spotUserService services.SpotUserService
	userService     services.UserService
}

func NewSpotUserHandler(spotUserService services.SpotUserService, userService services.UserService) *SpotUserHandler {
	return &SpotUserHandler{
		spotUserService: spotUserService,
		userService:     userService,
	}
}

func (su *SpotUserHandler) CreateSpotUser(c *gin.Context) {

	var requestBody CreateSpotUserRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by badge ID
	user, err := su.userService.GetUserByBadgeId(requestBody.BadgeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find user"})
		return
	}
	exSpotUser, err := su.spotUserService.GetActiveSpotUserByID(uint(user.ID))
	if exSpotUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already assigned to spot : " + exSpotUser.Spot.Name})
		return
	}
	spotId, err := strconv.ParseUint(requestBody.SpotID, 10, 64)
	// Create spot user
	spotUser := models.SpotUser{
		SpotID:    uint(spotId),
		UserID:    user.ID,
		InTime:    time.Now(),
		OutTime:   requestBody.OutTime,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = su.spotUserService.CreateSpotUser(&spotUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create spot user please check if user is already assigne to another spot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Spot user created successfully"})
}

func (su *SpotUserHandler) GetActiveSpotUserCount(c *gin.Context) {
	// Get all active spot users with the count grouped by spot_id
	activeSpotUsers, err := su.spotUserService.GetActiveSpotUserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active spot users"})
		return
	}

	c.JSON(http.StatusOK, activeSpotUsers)
}

func (su *SpotUserHandler) GetActiveUsersBySpotID(c *gin.Context) {
	spotIDstr := c.Param("spotId")
	spotID, err := strconv.ParseUint(spotIDstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spot ID"})
		return
	}

	users, err := su.spotUserService.GetActiveUsersBySpotID(uint(spotID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (su *SpotUserHandler) DeleteUserFromSpot(c *gin.Context) {
	spotUserID := c.Param("spotUserID")
	userID, err := strconv.ParseUint(spotUserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	// Retrieve active spot user from the database

	spotUser, err := su.spotUserService.GetActiveSpotUserByID(uint(userID))

	// Update spot user fields
	spotUser.Active = false
	spotUser.OutTime = time.Now()

	// Save the changes to the database
	if err := su.spotUserService.UpdateUser(spotUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user from spot"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted from spot successfully"})
}
