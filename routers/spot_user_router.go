package routers

import (
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/gin-gonic/gin"
)

type SpotUserRouter struct {
	Handler *handlers.SpotUserHandler
}

func (sur *SpotUserRouter) SetupRoutes(router *gin.Engine) {
	spotUserGroup := router.Group("/spot-users")
	{
		spotUserGroup.POST("", sur.Handler.CreateSpotUser)
		spotUserGroup.GET("/active-count", sur.Handler.GetActiveSpotUserCount)
		spotUserGroup.GET("/active-users/:spotId", sur.Handler.GetActiveUsersBySpotID)
		spotUserGroup.PUT("/user/:spotUserID", sur.Handler.DeleteUserFromSpot)
	}

}
