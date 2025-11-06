package repository

import (
    "baton/backend/internal/database"
    // "baton/internal/models"
)

func CreateUser(telegramID int64, username string) error {
    _, err := database.DB.Exec("INSERT OR IGNORE INTO users (telegram_id, username) VALUES (?, ?)", telegramID, username)
    return err
}

func GetTelegramIDByUsername(username string) (int64, error) {
    if username == "" {
        return 0, nil
    }
    var id int64
    err := database.DB.QueryRow(
        `SELECT telegram_id FROM users WHERE LOWER(username)=LOWER(?) LIMIT 1`,
        username,
    ).Scan(&id)
    // if err == sql.ErrNoRows {
    //     return 0, nil
    // }
    if err != nil {
        return 0, err
    }
    return id, nil
}