package character

import "gocionics/internal/entities/character"

type ICharacterRepository interface {
	GetByID(id int) (*character.Character, error)
	ListAll() (*[]character.Character, error)
}
