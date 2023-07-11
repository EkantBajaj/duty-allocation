package routers

import (
	"github.com/ekantbajaj/duty-allocation/db"
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/ekantbajaj/duty-allocation/middlewares"
	"github.com/ekantbajaj/duty-allocation/repositories"
	"github.com/ekantbajaj/duty-allocation/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.LoggingMiddleware())
	router.Use(cors.Default())
	// create Repository
	dbOrm := db.GetDB()
	spotRepository := repositories.NewSpotRepository(dbOrm)
	userRepository := repositories.NewUserRepository(dbOrm)
	spotUserRepository := repositories.NewSpotUserRepository(dbOrm)

	// Create services
	spotService := services.NewSpotService(spotRepository)
	userService := services.NewUserService(userRepository)
	spotUserService := services.NewSpotUserService(spotUserRepository)

	//Create Handlers
	userHandler := handlers.NewUserHandler(userService)
	spotHandler := handlers.NewSpotHandler(spotService)
	spotUserHandler := handlers.NewSpotUserHandler(spotUserService, userService)
	// Spot routes
	spotRouter := SpotRouter{Handler: spotHandler}
	spotRouter.SetupRoutes(router)

	// User routes
	userRouter := UserRouter{
		Handler: userHandler,
	}
	userRouter.SetupRoutes(router)

	// SpotUser routes

	spotUserRouter := SpotUserRouter{
		Handler: spotUserHandler,
	}
	spotUserRouter.SetupRoutes(router)

	// Add more routers for other functionalities

	return router
}
