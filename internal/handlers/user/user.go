package user

import (
	"github.com/gin-gonic/gin"
	usecase "gocionics/internal/usecases/user"
	"net/http"
)

type UserHandler struct {
	userUC *usecase.UserUseCase
}

func NewUserHandler(userUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) AssignCharacter(c *gin.Context) {
	userID := c.Param("id")
	characterID := c.Param("characterID")
	err := h.userUC.AssignCharacter(userID, characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	panic("implement me")

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	panic("implement me")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// soft delete
	panic("implement me")
}
