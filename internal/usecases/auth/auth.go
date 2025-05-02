package auth

import (
	"errors"
	"time"

	jwtT "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"gocionics/internal/entities"
	user_repo "gocionics/internal/repositories/user"
)

type AuthUseCase struct {
	userRepo  user_repo.IUserRepository
	secretKey string
}

func NewAuthUseCase(userRepo user_repo.IUserRepository, secretKey string) *AuthUseCase {
	return &AuthUseCase{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
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

func (uc *AuthUseCase) GenerateToken(user *entities.User) (string, error) {
	claims := jwtT.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwtT.NewWithClaims(jwtT.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.secretKey))
}

func (uc *AuthUseCase) GetUserByToken(tokenString string) (*entities.User, error) {
	token, err := jwtT.Parse(tokenString, func(token *jwtT.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtT.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(uc.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwtT.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return uc.userRepo.GetByID(int(userID))
		}
		return nil, errors.New("invalid user_id in token")
	}

	return nil, errors.New("invalid token")
}
