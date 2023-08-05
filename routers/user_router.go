package routers

import (
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Handler *handlers.UserHandler
}

func (ur *UserRouter) SetupRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("", ur.Handler.CreateUser)
		userRouter.POST("/login", ur.Handler.LoginUser)
		userRouter.GET("/:id", ur.Handler.GetUserByID)
		// Add more routes as needed
	}

}
