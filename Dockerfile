FROM golang:1.23-alpine AS builder

# Устанавливаем git, так как go mod может понадобиться прямое скачивание репозиториев
RUN apk update && apk add --no-cache git

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY . .
RUN go build -o /go/bin/user-service ./cmd/user-service

# Минимизируем образ
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/user-service /usr/local/bin/user-service

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/user-service"]
