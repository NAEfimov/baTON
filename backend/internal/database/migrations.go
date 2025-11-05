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
        `CREATE TABLE IF NOT EXISTS vacancies (
            telegram_id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT,
            skills TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE TABLE IF NOT EXISTS candidates (
            telegram_id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT,
            name TEXT,
            skills TEXT, 
            matching_score REAL,
            years INTEGER,
            education TEXT,
            experience TEXT,
            location TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE INDEX IF NOT EXISTS idx_vacancies_telegram ON vacancies(telegram_id);`,
        `CREATE INDEX IF NOT EXISTS idx_candidates_telegram ON candidates(telegram_id);`,
    }

    for _, query := range queries {
        _, err := DB.Exec(query)
        if err != nil {
            log.Fatal("Migration failed:", err)
        }
    }

    log.Println("Migrations completed")
}