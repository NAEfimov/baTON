package repository

import (
    "baton/internal/database"
    // "baton/internal/models"
)

func CreateUser(telegramID int64, username string) error {
    _, err := database.DB.Exec("INSERT OR IGNORE INTO users (telegram_id, username) VALUES (?, ?)", telegramID, username)
    return err
}