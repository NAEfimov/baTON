package main

import (
	"baton/backend/internal/database"
	"baton/backend/handlers"
	"log"
	"net/http"
	"os"
    "path/filepath"
)

var botToken string

func main() {
	log.Println("Starting application...")

    botToken := os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("FATAL: BOT_TOKEN environment variable not set. Check your .env file.")
    }
    dbDir := "/backend/data"
    dbPath := filepath.Join(dbDir, "baTON.db")

    if err := os.MkdirAll(dbDir, 0755); err != nil {
        log.Fatalf("Failed to create data directory: %v", err)
    }
	log.Println("Initializing database...")
    database.InitDB(dbPath)
	database.RunMigrations()

	handlers.SetupRoutes()
    log.Println("Server listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
