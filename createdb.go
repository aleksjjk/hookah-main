package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Подключаемся к базе данных (или создаем файл, если его нет)
    db, err := sql.Open("sqlite3", "hookah_bot.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Создаем таблицу users
    createTableQuery := `CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY
    );`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatal("Ошибка при создании таблицы:", err)
    } else {
        log.Println("Таблица users успешно создана.")
    }
}