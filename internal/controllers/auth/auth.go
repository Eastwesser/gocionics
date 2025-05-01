package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	authUC AuthUseCase
}

// @Summary User login
// @Description Authenticate user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginRequest true "Credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	password := ctx.PostForm("password")
	fmt.Println(password) // XDDD
}

// SetupRoutes registers auth routes
func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.POST("/auth/login", c.Login)
	// Add other auth routes
}
