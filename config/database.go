package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database credentials from .env
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Connect to the database
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("✅ Connected to the database successfully!")
}

// InitializeDatabase creates necessary tables if they don't exist
func InitializeDatabase() error {
	// Create users table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	fmt.Println("✅ Database tables initialized successfully!")
	return nil
}
