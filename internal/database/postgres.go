package database

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
    "os"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
    dbHost := os.Getenv("DB_HOST")     // db
    dbPort := os.Getenv("DB_PORT")     // 5432
    dbUser := os.Getenv("DB_USER")     // postgres
    dbPass := os.Getenv("DB_PASSWORD") // secret
    dbName := os.Getenv("DB_NAME")     // users_db

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        dbUser, dbPass, dbHost, dbPort, dbName)

    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        return nil, err
    }

    // Проверка соединения
    err = pool.Ping(context.Background())
    if err != nil {
        return nil, err
    }

    return pool, nil
}
