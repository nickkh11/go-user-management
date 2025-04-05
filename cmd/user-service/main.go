package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "sync"

    "google.golang.org/grpc"
    "github.com/nickkh11/go-user-management/internal/pb/user"
    "github.com/nickkh11/go-user-management/internal/database"
    "github.com/nickkh11/go-user-management/internal/services"
)

func main() {
    // Запускаем подключение к БД
    dbPool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("Cannot connect to DB: %v", err)
    }
    defer dbPool.Close()

    // Запуск HTTP
    // (Можно вынести в функцию, но для примера — прямо здесь)
    httpPort := getEnv("HTTP_PORT", "8080")
    httpMux := http.NewServeMux()
    httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from user-service HTTP!")
    })
    httpMux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            fmt.Fprintln(w, "Список пользователей")
        case http.MethodPost:
            fmt.Fprintln(w, "Создаем пользователя")
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    httpServer := &http.Server{
        Addr:    ":" + httpPort,
        Handler: httpMux,
    }

    // Запуск gRPC
    // (Аналогично можно вынести в отдельную функцию)
    grpcPort := getEnv("GRPC_PORT", "50051")
    lis, err := net.Listen("tcp", ":"+grpcPort)
    if err != nil {
        log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
    }
    grpcServer := grpc.NewServer()
    userpb.RegisterUserServiceServer(
        grpcServer,
        &services.UserServiceServer{
            // сюда, например, можно передавать dbPool
        },
    )

    // Запускаем обе «службы» (HTTP и gRPC) параллельно:
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        log.Printf("HTTP server is running on port %s", httpPort)
        if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("HTTP Server failed: %v", err)
        }
    }()

    go func() {
        defer wg.Done()
        log.Printf("gRPC server is running on port %s", grpcPort)
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("gRPC Server failed: %v", err)
        }
    }()

    // «Задерживаемся» тут, пока обе горутины не завершатся
    wg.Wait()
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}
