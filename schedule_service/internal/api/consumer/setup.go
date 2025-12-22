package consumer

import (
	"context"

	kafkalib "github.com/Quszlet/libs/kafka"
	"github.com/Quszlet/schedule_service/internal/kafka"
	"github.com/Quszlet/schedule_service/internal/service"
)

// Runner describes consumer and its run function.
type Runner struct {
	Consumer *kafkalib.Consumer
	Run      func(ctx context.Context) error
}

// PrepareConsumers instantiates all Kafka consumers for the service.
func PrepareConsumers(s *service.Service, p *kafka.Producers) ([]Runner, error) {
	eventHandler := NewEventHandler(s, p)

	kafkaConsumer, err := kafkalib.NewConsumer(
		[]string{"kafka:9092"},
		"doctors-events",
		"",
	)
	if err != nil {
		return nil, err
	}

	return []Runner{
		{
			Consumer: kafkaConsumer,
			Run: func(ctx context.Context) error {
				return ConsumeKafkaEvents(ctx, kafkaConsumer, eventHandler)
			},
		},
	}, nil
}
