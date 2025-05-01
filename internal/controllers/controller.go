package controllers

import "gocionics/internal/usecases"

type AuthController struct {
	authUC usecases.AuthUseCase
}

func (c *AuthController) Login(username, password string) (string, error) {
	panic("implement me")
}
func (c *AuthController) Logout() error {
	panic("implement me")
}

func NewAuthController(authUC usecases.AuthUseCase) *AuthController {
	return &AuthController{authUC: authUC}
}
