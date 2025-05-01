package auth

import (
	"errors"
	"gocionics/internal/entities/user"
	user_repo "gocionics/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo user_repo.IUserRepository
}

func NewAuthUseCase(userRepo user_repo.IUserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) Register(email, password string) (*user.User, error) {
	// Проверяем, существует ли пользователь
	existingUser, err := uc.userRepo.GetByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Создаем пользователя
	newUser := &user.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	userID, err := uc.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = userID
	return newUser, nil
}

func (uc *AuthUseCase) Login(email, password string) (*user.User, error) {
	// Находим пользователя
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (uc *AuthUseCase) GetUserByID(id string) (*user.User, error) {
	return uc.userRepo.GetByID(id)
}
