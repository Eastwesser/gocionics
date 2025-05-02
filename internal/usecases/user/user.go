package user

import (
	"errors"
	"gocionics/internal/entities"
	charrepo "gocionics/internal/repositories/character"
	userrepo "gocionics/internal/repositories/user"
)

type UserUseCase struct {
	userRepo userrepo.IUserRepository
	charRepo charrepo.ICharacterRepository
}

func NewUserUseCase(userRepo userrepo.IUserRepository, charRepo charrepo.ICharacterRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		charRepo: charRepo,
	}
}

func (uc *UserUseCase) AssignCharacter(userID int, characterID int) error {
	// Проверяем существование характера
	_, err := uc.charRepo.GetByID(characterID)
	if err != nil {
		return errors.New("character not found")
	}

	// Назначаем характер пользователю
	return uc.userRepo.AssignCharacter(userID, characterID)
}

func (uc *UserUseCase) GetUserCharacter(userID int) (*entities.Character, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user.CharacterID == 0 {
		return nil, errors.New("user has no character assigned")
	}

	return uc.charRepo.GetByID(user.CharacterID)
}
