package controllers

import (
	"gocionics/internal/usecases/auth"
)

type AuthController struct {
	authUC auth.AuthUseCase
}

func (c *AuthController) Login(username, password string) (string, error) {
	panic("implement me")
}
func (c *AuthController) Logout() error {
	panic("implement me")
}

func NewAuthController(authUC auth.AuthUseCase) *AuthController {
	return &AuthController{authUC: authUC}
}
