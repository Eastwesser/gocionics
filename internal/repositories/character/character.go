package character

import (
	"database/sql"
	"gocionics/internal/entities/character"
)

type ICharacterRepository interface {
	GetByID(id int) (*character.Character, error)
	ListAll() (*[]character.Character, error)
}

type CharacterRepo struct {
	db *sql.DB
}

func NewCharacterRepo(db *sql.DB) *CharacterRepo {
	return &CharacterRepo{db: db}
}

func (r *CharacterRepo) GetByID(id int) (*character.Character, error) {
	panic("implement me")
}

func (r *CharacterRepo) ListAll() (*[]character.Character, error) {
	panic("implement me")
}

func (r *CharacterRepo) Save(character *character.Character) error {
	panic("implement me")
}

func (r *CharacterRepo) Delete(id int) error {
	panic("implement me")
}

func (r *CharacterRepo) Close() error {
	panic("implement me")
}

func (r *CharacterRepo) Connect() error {
	panic("implement me")
}

func (r *CharacterRepo) Disconnect() {
	panic("implement me")
}

func (r *CharacterRepo) Query(query string, args ...interface{}) (*[]character.Character, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var characters []character.Character
	for rows.Next() {
		var char character.Character
		if characters = append(characters, char); err != nil {
			return nil, err
		}
	}
	return &characters, nil
}
