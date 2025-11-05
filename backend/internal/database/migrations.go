package database

import (
    "log"
)

func RunMigrations() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            telegram_id INTEGER UNIQUE,
            username TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
    }

    for _, query := range queries {
        _, err := DB.Exec(query)
        if err != nil {
            log.Fatal("Migration failed:", err)
        }
    }

    log.Println("Migrations completed")
}