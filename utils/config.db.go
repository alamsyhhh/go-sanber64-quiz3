package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found, menggunakan environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL tidak diset di environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("koneksi database gagal: %v", err)
	}

	fmt.Println("✅ Koneksi ke database berhasil!")
	return db, nil
}
