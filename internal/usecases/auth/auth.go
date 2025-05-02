package auth

import (
	"errors"
	"gocionics/internal/entities"
	user_repo "gocionics/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo user_repo.IUserRepository
}

func NewAuthUseCase(userRepo user_repo.IUserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) Register(email, password string) (*entities.User, error) {
	existingUser, err := uc.userRepo.GetByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entities.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	id, err := uc.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = id
	return newUser, nil
}

func (uc *AuthUseCase) Login(email, password string) (*entities.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (uc *AuthUseCase) GetUserByID(id int) (*entities.User, error) {
	return uc.userRepo.GetByID(id)
}
