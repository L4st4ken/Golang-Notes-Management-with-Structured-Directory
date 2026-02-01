package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// buka koneksi database
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// ping dengan timeout (hindari EOF tanpa pesan)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Database berhasil terkoneksi")

	return db, nil
}
