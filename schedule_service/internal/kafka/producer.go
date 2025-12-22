package kafka

import (
	"log/slog"

	kafkalib "github.com/Quszlet/libs/kafka"
)

// Producers keeps all Kafka producers used by the service.
type Producers struct {
	ScheduleEvents *kafkalib.Producer
}

// InitKafkaProducers centralizes creation of all Kafka producers.
func InitKafkaProducers() (*Producers, func(), error) {
	scheduleEventsProducer, err := kafkalib.NewProducer([]string{"kafka:9092"}, "schedule-events")
	if err != nil {
		return nil, func() {}, err
	}

	cleanup := func() {
		if err := scheduleEventsProducer.Close(); err != nil {
			slog.Error("failed to close kafka producer", "error", err)
		}
	}

	return &Producers{
		ScheduleEvents: scheduleEventsProducer,
	}, cleanup, nil
}
