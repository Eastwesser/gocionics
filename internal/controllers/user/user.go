package user

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userUC UserUseCase
}

// @Summary Assign character to user
// @Description Assign character to user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param character_id path int true "Character ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /users/{id}/characters/{character_id} [post]
func (c *Controller) AssignCharacter(ctx *gin.Context) {
	// Implementation
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.POST("/users/:id/characters/:character_id", c.AssignCharacter)
	// Add other user routes
}
