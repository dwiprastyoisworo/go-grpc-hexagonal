package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func PostgresInit(cfg *App) (*sqlx.DB, error) {
	// Format DSN (Data Source Name)
	log.Println("Connecting to PostgreSQL...")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	// Koneksi ke database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	db.SetConnMaxLifetime(time.Duration(cfg.Postgres.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Duration(cfg.Postgres.MaxIdleTime))
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConn)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConn)

	log.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
