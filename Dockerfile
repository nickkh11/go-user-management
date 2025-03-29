# Указываем базовый образ
FROM golang:1.20-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарник
RUN go build -o /go/bin/user-service ./cmd/user-service

# Минимизируем образ – используем базовый образ alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Копируем бинарник из предыдущего шага
COPY --from=builder /go/bin/user-service /usr/local/bin/user-service

# Говорим, на каком порту будет работать наше приложение
EXPOSE 8080

# Запускаем бинарник
ENTRYPOINT ["/usr/local/bin/user-service"]
