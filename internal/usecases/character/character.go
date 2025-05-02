package character

import (
	"errors"
	"gocionics/internal/entities"
	charrepo "gocionics/internal/repositories/character"
)

type CharacterUseCase struct {
	repo charrepo.ICharacterRepository
}

func NewCharacterUseCase(repo charrepo.ICharacterRepository) *CharacterUseCase {
	return &CharacterUseCase{repo: repo}
}

func (uc *CharacterUseCase) GetByID(id int) (*entities.Character, error) {
	return uc.repo.GetByID(id)
}

func (uc *CharacterUseCase) ListAll() ([]*entities.Character, error) {
	return uc.repo.ListAll()
}

func (uc *CharacterUseCase) AnalyzeAnswers(answers []int) (*entities.Character, error) {
	if len(answers) == 0 {
		return nil, errors.New("no answers provided")
	}

	// Простейшая логика - выбираем тип по количеству ответов
	characters, err := uc.repo.ListAll()
	if err != nil || len(characters) == 0 {
		return nil, errors.New("no characters available")
	}

	index := len(answers) % len(characters)
	return characters[index], nil
}
