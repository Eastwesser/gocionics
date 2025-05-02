package server

import (
	"github.com/gin-gonic/gin"
	authcontroller "gocionics/internal/controllers/auth"
	charactercontroller "gocionics/internal/controllers/character"
	usercontroller "gocionics/internal/controllers/user"
	"gocionics/internal/middleware"
	"time"
)

func SetupRoutes(router *gin.Engine,
	authCtrl *authcontroller.Controller,
	userCtrl *usercontroller.Controller,
	charCtrl *charactercontroller.Controller) {

	// Глобальные middleware
	router.Use(middleware.RateLimiter(100, time.Minute)) // 100 запросов в минуту

	api := router.Group("/api/v1")
	{
		authcontroller.SetupRoutes(api, authCtrl)
		usercontroller.SetupRoutes(api, userCtrl)
		charactercontroller.SetupRoutes(api, charCtrl)
	}
}
