package db

import (
	"database/sql"
	"gocionics/internal/entities/user"
)

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) Init() {
	//...
}

func (r *UserRepo) Close() {
	// ...
}

func (r *UserRepo) Create(user *user.User) error {
	_, err := r.db.Exec("INSERT INTO users...")
	return err
}
