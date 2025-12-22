package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	group  string
}

type MessageHandler func(ctx context.Context, event *Event) error

func NewConsumer(brokers []string, topic, groupID string) (*Consumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupID,
		StartOffset:    kafka.LastOffset,
		CommitInterval: 0, // disable auto-commit
	})

	slog.Info("Kafka consumer initialized", "brokers", brokers, "topic", topic, "group", groupID)
	return &Consumer{reader: reader, group: groupID}, nil
}

func (c *Consumer) StartConsuming(ctx context.Context, handler MessageHandler) error {
	slog.Info("Starting to consume messages from Kafka", "topic", c.reader.Config().Topic)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Stopping Kafka consumer")
			return ctx.Err()
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				slog.Error("failed to read message from kafka", "error", err.Error())
				continue
			}

			var event Event
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				slog.Error("failed to unmarshal kafka message", "error", err.Error())
				continue
			}

			slog.Debug("received kafka message", "event_type", event.EventType)

			if err := handler(ctx, &event); err != nil {
				slog.Error("failed to handle kafka event", "error", err.Error(), "event_type", event.EventType)
				continue
			}

			if c.group != "" {
				if err := c.reader.CommitMessages(ctx, msg); err != nil {
					slog.Error("failed to commit message", "error", err.Error())
				}
			}
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
