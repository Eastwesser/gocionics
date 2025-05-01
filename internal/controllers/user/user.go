package user

import (
	"github.com/gin-gonic/gin"
	"gocionics/internal/usecases/user"
	"net/http"
	"strconv"
)

type Controller struct {
	userUC *user.UserUseCase
}

func NewUserController(userUC *user.UserUseCase) *Controller {
	return &Controller{userUC: userUC}
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
	userID := ctx.Param("id")
	characterID, err := strconv.Atoi(ctx.Param("character_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	if err := c.userUC.AssignCharacter(userID, characterID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "character assigned"})
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.POST("/users/:id/characters/:character_id", c.AssignCharacter)
}
