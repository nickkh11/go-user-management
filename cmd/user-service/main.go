package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/nickkh11/go-user-management/internal/database"
)

func main() {
    port := "8080"
    if fromEnv := os.Getenv("PORT"); fromEnv != "" {
        port = fromEnv
    }

    // Создаем подключение к БД
    dbPool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("Cannot connect to DB: %v", err)
    }
    defer dbPool.Close()

    // Пример обработки GET / для проверки
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from user-service!")
    })

    // Пример хендлера для получения списка пользователей
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            // Логика чтения пользователей из dbPool
            fmt.Fprintln(w, "Список пользователей")
        case http.MethodPost:
            // Логика создания нового пользователя
            fmt.Fprintln(w, "Создаем пользователя")
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Printf("Starting server on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
