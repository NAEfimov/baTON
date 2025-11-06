package models

type Candidate struct {
    TelegramID    int64            `json:"telegram_id"`
    Username      string           `json:"username"`
    Name          string           `json:"name"`
    Skills        []string         `json:"skills"`
    MatchingScore float64          `json:"matching_score"`
    Years         int              `json:"years"`
    Education     string           `json:"education"`
    Experience    []ExperienceItem `json:"experience"`
    Location      string           `json:"location"`
}

type ExperienceItem struct {
    Company    string   `json:"company"`
    Vacancy    string   `json:"vacancy"`
    Highlights []string `json:"highlights"`
    Verified   bool     `json:"verified"`
}

type CandidatePublicDTO struct {
    Username      string           `json:"username"`
    Name          string           `json:"name"`
    MatchingScore float64          `json:"matching_score"`
    Years         int              `json:"years"`
    Education     string           `json:"education"`
    Experience    []ExperienceItem `json:"experience"`
    Location      string           `json:"location"`
}