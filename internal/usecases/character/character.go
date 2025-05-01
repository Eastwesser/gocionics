package character

import "gocionics/internal/repositories/character"

type CharacterUseCase struct {
	repo character.ICharacterRepository
}

func NewCharacterUseCase(repo character.ICharacterRepository) *CharacterUseCase {
	return &CharacterUseCase{repo: repo}
}

func (uc *CharacterUseCase) GetCharacter(id int) (*character.Character, error) {
	return uc.repo.GetByID(id)
}
