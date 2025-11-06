package models

type Vacancy struct {
    TelegramID    int64            `json:"telegram_id"`
    Username      string           `json:"username"`
    Skills        []string         `json:"skills"`
}