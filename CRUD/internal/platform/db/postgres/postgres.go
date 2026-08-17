package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	dsn := "postgresql://postgres:Shishir123@localhost:5433/go_crud?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("PostgreSQL connected")

	return db, nil
}