package routers

import (
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/gin-gonic/gin"
)

type SpotRouter struct {
	Handler *handlers.SpotHandler
}

func (sr *SpotRouter) SetupRoutes(router *gin.Engine) {
	spotGroup := router.Group("/spots")
	{
		spotGroup.POST("/", sr.Handler.CreateSpot)
		spotGroup.POST("", sr.Handler.CreateSpot)
		spotGroup.GET("/:id", sr.Handler.GetSpot)
		spotGroup.GET("", sr.Handler.GetAllSpots)
		// Add more spot routes as needed
	}
}
