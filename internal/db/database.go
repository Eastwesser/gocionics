package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gocionics/config"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_host, cfg.DB_port, cfg.DB_user, cfg.DB_password, cfg.DB_name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}
