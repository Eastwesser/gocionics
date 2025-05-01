package character

import (
	"errors"
	"gocionics/internal/entities/character"
	char_repo "gocionics/internal/repositories/character"
)

type CharacterUseCase struct {
	repo char_repo.ICharacterRepository
}

func NewCharacterUseCase(repo char_repo.ICharacterRepository) *CharacterUseCase {
	return &CharacterUseCase{repo: repo}
}

func (uc *CharacterUseCase) GetByID(id int) (*character.Character, error) {
	return uc.repo.GetByID(id)
}

func (uc *CharacterUseCase) ListAll() ([]*character.Character, error) {
	return uc.repo.ListAll()
}

func (uc *CharacterUseCase) AnalyzeAnswers(answers []int) (*character.Character, error) {
	// Здесь должна быть логика анализа ответов на тест
	// Пока возвращаем первый попавшийся характер
	characters, err := uc.repo.ListAll()
	if err != nil || len(characters) == 0 {
		return nil, errors.New("no characters available")
	}
	return characters[0], nil
}
