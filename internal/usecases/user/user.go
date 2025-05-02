package user

import (
	"errors"
	"gocionics/internal/entities"
	char_repo "gocionics/internal/repositories/character"
	user_repo "gocionics/internal/repositories/user"
)

type UserUseCase struct {
	userRepo user_repo.IUserRepository
	charRepo char_repo.ICharacterRepository
}

func NewUserUseCase(userRepo user_repo.IUserRepository, charRepo char_repo.ICharacterRepository) *UserUseCase {
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
