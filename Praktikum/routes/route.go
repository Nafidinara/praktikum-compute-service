package routes

import (
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Praktikum/controllers"
	"Praktikum/middlewares"
)

func InitRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	e.Use(middleware.Recover())

	rateLimiterConfig := middlewares.RateLimiterConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 10 * time.Minute,
	}

	rateLimiterMiddleware := rateLimiterConfig.Init()

	e.Use(rateLimiterMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       "itoldyouthisissosecret",
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserController(&jwtConfig)

	auth := e.Group("/api/v1/auth")

	auth.POST("/login", userController.Login)
	auth.POST("/register", userController.Register)

	users := e.Group("/api/v1/users", echojwt.WithConfig(authMiddlewareConfig))
	users.GET("/me", userController.GetUser)

	packages := e.Group("/api/v1/packages", echojwt.WithConfig(authMiddlewareConfig))
	packages.GET("/", controllers.GetAllPackage)
	packages.POST("/", controllers.CreatePackage)
	packages.GET("/:id", controllers.GetOnePackage)
	packages.PUT("/:id", controllers.UpdatePackage)
	packages.DELETE("/:id", controllers.DeletePackage)
}
