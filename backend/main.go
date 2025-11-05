package main

import (
	"baton/backend/internal/database"
	"baton/backend/internal/telegram"
	"log"
	"net/http"
	"os"
)

var botToken string

func main() {
	log.Println("Starting application...")
	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN required")
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
