package auth

import (
	"github.com/gin-gonic/gin"
	"gocionics/internal/usecases/auth"
	"net/http"
)

type Controller struct {
	authUC *auth.AuthUseCase
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewAuthController(authUC *auth.AuthUseCase) *Controller {
	return &Controller{authUC: authUC}
}

func (c *Controller) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authUC.Register(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
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
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authUC.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.POST("/register", c.Register)
	r.POST("/login", c.Login)
}
