package character

import (
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} Character
// @Failure 404 {object} ErrorResponse
// @Router /characters/{id} [get]
func (c *Controller) GetCharacter(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	char, err := c.charUC.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, char)
}

func (c *Controller) ListCharacters(ctx *gin.Context) {
	characters, err := c.charUC.ListAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, characters)
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.GET("/characters", c.ListCharacters)
	r.GET("/characters/:id", c.GetCharacter)
}
