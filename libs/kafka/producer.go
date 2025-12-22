package kafka

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

// Producer wraps kafka.Writer to emit domain events.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer configures writer for brokers/topic.
func NewProducer(brokers []string, topic string) (*Producer, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	slog.Info("Kafka producer initialized", "brokers", brokers, "topic", topic)
	return &Producer{writer: writer}, nil
}

// SendMessage serializes event data and pushes to Kafka.
func (p *Producer) SendMessage(ctx context.Context, eventType string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal kafka payload", "error", err.Error())
		return err
	}

	event := Event{
		EventType: eventType,
		Timestamp: time.Now().UnixMilli(),
		Data:      payload,
	}

	value, err := json.Marshal(event)
	if err != nil {
		slog.Error("failed to marshal kafka event", "error", err.Error())
		return err
	}

	if err := p.writer.WriteMessages(ctx, kafka.Message{Value: value}); err != nil {
		slog.Error("failed to write message to kafka", "error", err.Error())
		return err
	}

	slog.Debug("message sent to kafka", "event_type", eventType)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
