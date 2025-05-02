package character

import (
	"github.com/gin-gonic/gin"
	"gocionics/internal/entities"
	"gocionics/internal/usecases/character"
	"net/http"
	"strconv"
)

type Controller struct {
	charUC *character.CharacterUseCase
}

func NewCharacterController(charUC *character.CharacterUseCase) *Controller {
	return &Controller{charUC: charUC}
}

// @Summary Get character by ID
// @Description Get character details
// @Tags character
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} entities.Character
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Router /characters/{id} [get]
func (c *Controller) GetCharacter(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entities.ErrorResponse{Error: "invalid character ID"})
		return
	}

	char, err := c.charUC.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, char)
}

// @Summary Get all characters
// @Description Get list of all available character types
// @Tags character
// @Produce json
// @Success 200 {array} entities.Character
// @Failure 500 {object} entities.ErrorResponse
// @Router /characters [get]
func (c *Controller) ListCharacters(ctx *gin.Context) {
	characters, err := c.charUC.ListAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entities.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, characters)
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.GET("/characters", c.ListCharacters)
	r.GET("/characters/:id", c.GetCharacter)
}
