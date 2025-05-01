package user

import (
	"database/sql"
	"errors"
	"gocionics/internal/entities/user"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(u *user.User) (string, error) {
	query := `
		INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) 
		RETURNING id`

	var id string
	err := r.db.QueryRow(query, u.Email, u.Password).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *PostgresRepository) GetByID(id string) (*user.User, error) {
	query := `
		SELECT id, email, password_hash, character_id 
		FROM users 
		WHERE id = $1`

	var u user.User
	err := r.db.QueryRow(query, id).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.CharacterID,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresRepository) AssignCharacter(userID string, characterID int) error {
	query := `
		UPDATE users 
		SET character_id = $1 
		WHERE id = $2`

	_, err := r.db.Exec(query, characterID, userID)
	return err
}

func (r *PostgresRepository) GetByEmail(email string) (*user.User, error) {
	query := `
		SELECT id, email, password_hash 
		FROM users 
		WHERE email = $1`

	var u user.User
	err := r.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}
