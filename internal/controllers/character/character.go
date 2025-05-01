package character

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	charUC CharacterUseCase
}

// @Summary Get character by ID
// @Description Get character details
// @Tags character
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} Character
// @Failure 404 {object} ErrorResponse
// @Router /characters/{id} [get]
func (c *Controller) GetCharacter(ctx *gin.Context) {
	// Implementation
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.GET("/characters/:id", c.GetCharacter)
	// Add other character routes
}
