package handlers

import (
    "baton/backend/internal/models"
    "baton/backend/internal/repository"
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "sort"
    "strconv"
    "strings"
)

func HandleCreateCandidateJSON(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var p models.Candidate
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        log.Printf("Decode error: %v", err)
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    p.Skills = NormalizeSkills(p.Skills)
    var expBuf bytes.Buffer
    if err := json.NewEncoder(&expBuf).Encode(p.Experience); err != nil {
        log.Printf("Encode experience error: %v", err)
        http.Error(w, "encode experience error", http.StatusInternalServerError)
        return
    }
    _, err := repository.CreateCandidate(
        p.TelegramID,
        p.Username,
        p.Name,
        p.Skills,
        p.MatchingScore,
        p.Years,
        p.Education,
        strings.TrimSpace(expBuf.String()),
        p.Location,
    )
    if err != nil {
        log.Printf("CreateCandidate DB error: %v", err) 
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(map[string]any{
        "ok":       true,
        "username": p.Username,
    })
}

func NormalizeSkills(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(strings.ToLower(s))
		if s == "" {
			continue
		}
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func HandleGetCandidate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    telegramIDStr := r.URL.Query().Get("telegram_id")
    if telegramIDStr == "" {
        http.Error(w, "telegram_id required", http.StatusBadRequest)
        return
    }
    telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
    if err != nil {
        http.Error(w, "invalid telegram_id", http.StatusBadRequest)
        return
    }
    candidate, err := repository.GetCandidateByTelegramID(telegramID)
    if err != nil {
        log.Printf("GetCandidateByTelegramID error: %v", err)
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }
    if candidate == nil {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    dto := models.CandidatePublicDTO{
        Username:      candidate.Username,
        Name:          candidate.Name,
        MatchingScore: candidate.MatchingScore,
        Years:         candidate.Years,
        Education:     candidate.Education,
        Experience:    candidate.Experience,
        Location:      candidate.Location,
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(dto)
}

func HandleGetMatchedCandidates(w http.ResponseWriter, r *http.Request) {
    telegramIDStr := r.URL.Query().Get("telegram_id")
    if telegramIDStr == "" {
        http.Error(w, "telegram_id required", http.StatusBadRequest)
        return
    }
    telegramID, _ := strconv.ParseInt(telegramIDStr, 10, 64)
    vacancy, _ := repository.GetVacancyByTelegramID(telegramID)
    if vacancy == nil {
        http.Error(w, "vacancy not found", http.StatusNotFound)
        return
    }
    candidates, _ := repository.GetAllCandidates()
    type Match struct {
        Candidate *models.Candidate
        Score     float64
    }
    var matches []Match
    vacancySkills := make(map[string]bool)
    for _, s := range vacancy.Skills {
        vacancySkills[s] = true
    }
    for _, c := range candidates {
        matchCount := 0
        for _, s := range c.Skills {
            if vacancySkills[s] {
                matchCount++
            }
        }
        score := float64(matchCount) / float64(len(vacancy.Skills)) * 100
        matches = append(matches, Match{c, score})
    }
    sort.Slice(matches, func(i, j int) bool {
        return matches[i].Score > matches[j].Score
    })
    if len(matches) > 5 {
        matches = matches[:5]
    }
    result := []map[string]any{}
    for _, m := range matches {
        result = append(result, map[string]any{
            "name":           m.Candidate.Name,
            "username":       m.Candidate.Username,
            "matching_score": m.Score,
            "years":          m.Candidate.Years,
            "education":      m.Candidate.Education,
            "experience":     m.Candidate.Experience,
            "location":       m.Candidate.Location,
        })
    }
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(result)
}