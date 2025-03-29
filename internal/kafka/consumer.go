package kafka

import (
    "context"
    "os"
    "github.com/segmentio/kafka-go"
)

func NewKafkaReader() *kafka.Reader {
    brokerAddress := os.Getenv("KAFKA_BROKER") // kafka:9092
    return kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{brokerAddress},
        GroupID: "user-service-group",
        Topic:   "user-events",
    })
}

// Чтение сообщений в отдельной горутине
func ConsumeMessages(reader *kafka.Reader) {
    go func() {
        defer reader.Close()
        for {
    _, err := reader.ReadMessage(context.Background())
            if err != nil {
                // Ошибка чтения
                break
            }
            // Обработка сообщения
            // m.Key и m.Value – это []byte
        }
    }()
}
