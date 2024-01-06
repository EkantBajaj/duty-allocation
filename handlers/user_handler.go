package handlers

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"github.com/ekantbajaj/duty-allocation/services"
	"github.com/ekantbajaj/duty-allocation/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var err error
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	err = uh.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	// Convert userID to uint
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the GetUserByID method of the user service
	user, err := uh.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// Return the user details
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) LoginUser(c *gin.Context) {
	// Get the badge ID and password from the request body
	var user *models.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// get user by badge id
	dbuser, err := uh.userService.GetUserByBadgeId(user.BadgeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found"})
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received

	if err := util.CheckPassword(dbuser.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate token through tokenMaker in repository
	token, err := uh.userService.CreateToken(dbuser.BadgeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}
	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})

}

// Implement other user-related handler methods as needed
