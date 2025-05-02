package auth

import (
	"errors"
	"gocionics/internal/entities"
	user_repo "gocionics/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (uc *AuthUseCase) GetUserByToken(tokenString string) (*entities.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["user_id"].(float64))
		return uc.userRepo.GetByID(userID)
	}

	return nil, errors.New("invalid token")
}

func (uc *AuthUseCase) GenerateToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte("your-secret-key"))
}
