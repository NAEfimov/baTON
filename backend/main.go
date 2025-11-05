package main

import (
	"baton/backend/internal/database"
	"baton/backend/handlers"
	"bufio"
    // "encoding/json"
    // "io"
	"log"
	"net/http"
	"os"
	"strings"
)

var botToken string

func main() {
	log.Println("Starting application...")

	// Read and parse config.env
    file, err := os.Open("config.env")
    if err != nil {
        log.Fatal("Error opening config.env:", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(line, "BOT_TOKEN=") {
            botToken = strings.TrimPrefix(line, "BOT_TOKEN=")
            break
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("Error reading config.env:", err)
    }
    if botToken == "" {
        log.Fatal("BOT_TOKEN not found in config.env")
    }
	//Initialize database
	log.Println("Initializing database...")
	database.InitDB("./baTON.db")
	database.RunMigrations()
	handlers.SetupRoutes()
    log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
