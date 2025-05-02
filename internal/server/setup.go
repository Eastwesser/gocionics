package server

import (
	"github.com/gin-gonic/gin"
	authcontroller "gocionics/internal/controllers/auth"
	charactercontroller "gocionics/internal/controllers/character"
	"gocionics/internal/controllers/health"
	usercontroller "gocionics/internal/controllers/user"
	"gocionics/internal/middleware"
	"time"
)

func SetupRoutes(router *gin.Engine,
	authCtrl *authcontroller.Controller,
	userCtrl *usercontroller.Controller,
	charCtrl *charactercontroller.Controller) {

	healthCtrl := health.NewController()
	router.GET("/health", healthCtrl.Status)

	// Глобальные middleware
	router.Use(middleware.RateLimiter(100, time.Minute)) // 100 запросов в минуту

	// API v1
	api := router.Group("/api/v1")
	{
		authcontroller.SetupRoutes(api, authCtrl)
		usercontroller.SetupRoutes(api, userCtrl)
		charactercontroller.SetupRoutes(api, charCtrl)
	}
}
