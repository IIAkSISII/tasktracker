package database

import (
	"database/sql"
	"fmt"
	"github.com/IIAkSISII/tasktracker/internal/config"

	_ "github.com/lib/pq"
)

func Connect(ctg config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		ctg.Host, ctg.Port, ctg.User, ctg.Password, ctg.Name, ctg.SslMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
