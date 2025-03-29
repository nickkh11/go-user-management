package kafka

import (
    "context"
    "os"
    "time"

    "github.com/segmentio/kafka-go"
)

func NewKafkaWriter() *kafka.Writer {
    brokerAddress := os.Getenv("KAFKA_BROKER") // kafka:9092
    return &kafka.Writer{
        Addr:         kafka.TCP(brokerAddress),
        Topic:        "user-events",
        Balancer:     &kafka.LeastBytes{},
        RequiredAcks: kafka.RequireAll,
        Async:        false,
    }
}

// Пример отправки сообщения
func SendMessage(writer *kafka.Writer, key, value string) error {
    msg := kafka.Message{
        Key:   []byte(key),
        Value: []byte(value),
        Time:  time.Now(),
    }
    return writer.WriteMessages(context.Background(), msg)
}
