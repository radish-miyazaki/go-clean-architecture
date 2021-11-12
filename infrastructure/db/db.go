package db

import (
	"database/sql"
	"fmt"

	"github.com/radish-miyazaki/go-clean-architecture/config"
)

// NewDB データベースに接続する
func NewDB() (*sql.DB, error) {
	c := config.Config
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUsername,
		c.DBPassword,
		c.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// データベースの接続を待つ
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
