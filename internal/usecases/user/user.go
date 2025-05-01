package usecase

import (
	charRepo "gocionics/internal/repositories/character"
	userRepo "gocionics/internal/repositories/user"
)

type UserUseCase struct {
	userRepo userRepo.IUserRepository
	charRepo charRepo.ICharacterRepository
}

func NewUserUseCase(userRepo userRepo.IUserRepository, charRepo charRepo.ICharacterRepository) *UserUseCase {
	return &UserUseCase{userRepo, charRepo}
}

func (uc *UserUseCase) AssignCharacter(userID, characterID int) error {
	character, err := uc.charRepo.GetByID(characterID)
	if err != nil {
		return err
	}

	return character, uc.userRepo.AssignCharacter(userID, characterID)
}
