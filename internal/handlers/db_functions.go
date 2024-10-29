package handlers

import (
	"database/sql"
	"log"
)
func (h *Handler) addUserToDatabase(id string) error {
    db, err := sql.Open("sqlite3", "hookah_bot.db")
    if err != nil {
        log.Printf("Ошибка открытия базы данных: %v\n", err)
        return err
    }
    defer db.Close()

    exists, err := h.userExists(id)
    if err != nil {
        log.Printf("Ошибка проверки существования пользователя: %v\n", err)
        return err
    }

    // Если пользователя нет, добавляем его
    if !exists {
        _, err = db.Exec("INSERT INTO users (id) VALUES (?)", id)
        if err != nil {
            log.Printf("Ошибка вставки пользователя в базу данных: %v\n", err)
            return err
        }
        log.Println("Пользователь добавлен в базу данных.")
    } else {
        log.Println("Пользователь уже существует в базе данных.")
    }

    return nil
}

func (h *Handler) userExists(id string) (bool, error) {
    db, err := sql.Open("sqlite3", "hookah_bot.db")
    if err != nil {
        log.Printf("Ошибка открытия базы данных: %v\n", err)
        return false, err
    }
    defer db.Close()

    var exists bool
    err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists)
    if err != nil {
        log.Printf("Ошибка выполнения запроса EXISTS: %v\n", err)
        return false, err
    }

    return exists, nil
}