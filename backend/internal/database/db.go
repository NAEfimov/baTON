package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbPath string) {
    var err error
    DB, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatal("Failed to open database:", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }
    log.Println("Database connected")
}