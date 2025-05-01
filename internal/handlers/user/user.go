package user

import (
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

}

func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (handler *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	panic("implement me")

}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	panic("implement me")
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// soft delete
	panic("implement me")
}
