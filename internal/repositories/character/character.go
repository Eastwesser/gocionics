package character

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"gocionics/internal/entities/character"
	"strings"
)

type ICharacterRepository interface {
	GetByID(id int) (*character.Character, error)
	ListAll() ([]*character.Character, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetByID(id int) (*character.Character, error) {
	query := `SELECT id, type, description, traits FROM characters WHERE id = $1`
	var c character.Character
	err := r.db.QueryRow(query, id).Scan(
		&c.ID,
		&c.Type,
		&c.Description,
		pq.Array(&c.Traits),
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("character not found")
	}
	if err != nil {
		return nil, err
	}

	c.Traits = strings.Split(traitsStr, ",")
	return &c, nil
}

func (r *PostgresRepository) ListAll() ([]*character.Character, error) {
	query := `
		SELECT id, type, description, traits 
		FROM characters`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*character.Character
	for rows.Next() {
		var c character.Character
		var traitsStr string

		if err := rows.Scan(
			&c.ID,
			&c.Type,
			&c.Description,
			&traitsStr,
		); err != nil {
			return nil, err
		}

		c.Traits = strings.Split(traitsStr, ",")
		characters = append(characters, &c)
	}

	return characters, nil
}
