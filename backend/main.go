package main

import (
	"baton/backend/internal/database"
	"baton/backend/internal/telegram"
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
	database.InitDB("./baTON.db") // Creates baTON.db in project root
	log.Println("Initializing database...")
	database.RunMigrations()
	log.Println("Initializing database...")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/webhook", telegram.HandleWebhook)
	http.HandleFunc("/init/verify", telegram.HandleVerify)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
