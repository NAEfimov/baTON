package main

import (
	"baton/internal/database"
	"baton/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var botToken string

func main() {
	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN required")
	}

	// Initialize database
    database.InitDB("./baTON.db") // Creates baTON.db in project root
    database.RunMigrations()

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/init/verify", handleVerify)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var update struct {
        Message struct {
            From struct {
                ID       int64  `json:"id"`
                Username string `json:"username"`
            } `json:"from"`
        } `json:"message"`
    }
    if err := json.Unmarshal(body, &update); err != nil {
        log.Printf("Failed to unmarshal: %v", err)
        w.WriteHeader(200)
        return
    }

    if update.Message.From.ID != 0 {
        if err := repository.CreateUser(update.Message.From.ID, update.Message.From.Username); err != nil {
            log.Printf("Failed to save user: %v", err)
        }
    }
	
	log.Printf("update: %s\n", string(body))
	w.WriteHeader(200)
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InitData string `json:"initData"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	resp := map[string]interface{}{
		"ok":       true,
		"initData": req.InitData,
	}
	json.NewEncoder(w).Encode(resp)
}
