package routers

import (
	"github.com/ekantbajaj/duty-allocation/db"
	"github.com/ekantbajaj/duty-allocation/handlers"
	"github.com/ekantbajaj/duty-allocation/middlewares"
	"github.com/ekantbajaj/duty-allocation/repositories"
	"github.com/ekantbajaj/duty-allocation/services"
	"github.com/ekantbajaj/duty-allocation/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func SetupRouter() *gin.Engine {
	tokenMaker, _ := token.NewPasetoMaker(viper.GetString("token.secret"))
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change this to restrict allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middlewares.LoggingMiddleware())
	// router.Use(middlewares.AuthMiddleware(tokenMaker))
	// create Repository
	dbOrm := db.GetDB()
	spotRepository := repositories.NewSpotRepository(dbOrm)
	userRepository := repositories.NewUserRepository(dbOrm, tokenMaker)
	spotUserRepository := repositories.NewSpotUserRepository(dbOrm)
	gateRepository := repositories.NewGateRepository(dbOrm)

	// Create services
	spotService := services.NewSpotService(spotRepository)
	userService := services.NewUserService(userRepository)
	spotUserService := services.NewSpotUserService(spotUserRepository)
	gateService := services.NewGateService(gateRepository)

	//Create Handlers
	userHandler := handlers.NewUserHandler(userService)
	spotHandler := handlers.NewSpotHandler(spotService)
	spotUserHandler := handlers.NewSpotUserHandler(spotUserService, userService)
	gateHandler := handlers.NewGateHandler(gateService)
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

	// Gate routes
	gateRouter := GateRouter{
		Handler: gateHandler,
	}
	gateRouter.SetupRoutes(router)

	// Add more routers for other functionalities

	return router
}
