package handlers

import (
    "baton/backend/internal/telegram"
    "net/http"
)

func SetupRoutes() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    // Telegram webhook
    http.HandleFunc("/webhook", telegram.HandleWebhook)
    http.HandleFunc("/init/verify", telegram.HandleVerify)
    // Candidates
    http.HandleFunc("/candidates", candidatesHandler)
    // Vacancies
    http.HandleFunc("/vacancies", vacanciesHandler)
    //Matched candidate
    http.HandleFunc("/vacancies/matches", HandleGetMatchedCandidates)
}

func candidatesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        HandleCreateCandidateJSON(w, r)
    case http.MethodGet:
        HandleGetCandidate(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func vacanciesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        HandleCreateVacancyJSON(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
