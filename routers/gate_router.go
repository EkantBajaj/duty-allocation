package routers

import (
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/gin-gonic/gin"
)

type GateRouter struct {
	Handler *handlers.GateHandler
}

func (gr *GateRouter) SetupRoutes(router *gin.Engine) {
	gateGroup := router.Group("/gate")
	{
		gateGroup.POST("/entry", gr.Handler.CreateGateEntry)
		gateGroup.GET("/count", gr.Handler.GetGateCountByGate)
		gateGroup.GET("/total", gr.Handler.GetTotalGateCount)
		gateGroup.GET("/", gr.Handler.GetGates)
	}
}