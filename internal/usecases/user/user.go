package usecase

import "gocionics/internal/repositories/user"

type AuthUseCase struct {
	repo user.IUserRepository
}

func (uc *AuthUseCase) Authenticate(email string, password string) (string, error) {
	panic("implement me")
}

func (uc *AuthUseCase) Register(email string, password string) error {
	panic("implement me")
}

func (uc *AuthUseCase) Login(email string, password string) error {
	panic("implement me")
}

func (uc *AuthUseCase) Logout() error {

	panic("implement me")
}

func (uc *AuthUseCase) RefreshToken(email string, password string) error {
	panic("implement me")
}

func (uc *AuthUseCase) ResetPassword(email string, password string) error {
	panic("implement me")
}

func (uc *AuthUseCase) Reset(email string, password string) error {
	panic("implement me")
}

func NewUserUseCase() {

}
