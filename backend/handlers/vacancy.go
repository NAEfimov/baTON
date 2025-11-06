package handlers

import (
	"baton/backend/internal/models"
	"baton/backend/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

func HandleCreateVacancyJSON(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var p models.Vacancy
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        log.Printf("Decode error: %v", err)
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    p.Skills = NormalizeSkills(p.Skills)
  	_, err := repository.CreateVacancy(
        p.TelegramID,
        p.Username,
        p.Skills,
    )
    if err != nil {
        log.Printf("CreateVacancy DB error: %v", err)
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]any{
        "ok": true,
        "telegram_id": p.TelegramID,
        "username": p.Username,
    })
}

