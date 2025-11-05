package models

type User struct {
    ID         int    `json:"id"`
    TelegramID int64  `json:"telegram_id"`
    Username   string `json:"username"`
}