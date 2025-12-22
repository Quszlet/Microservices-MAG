package kafka

import (
	"log/slog"

	kafkalib "github.com/Quszlet/libs/kafka"
)

// Producers keeps all Kafka producers used by the service.
type Producers struct {
	BookingEvents *kafkalib.Producer
}

// InitKafkaProducers centralizes creation of all Kafka producers.
func InitKafkaProducers() (*Producers, func(), error) {
	bookingEventsProducer, err := kafkalib.NewProducer([]string{"kafka:9092"}, "booking-events")
	if err != nil {
		return nil, func() {}, err
	}

	cleanup := func() {
		if err := bookingEventsProducer.Close(); err != nil {
			slog.Error("failed to close kafka producer", "error", err)
		}
	}

	return &Producers{
		BookingEvents: bookingEventsProducer,
	}, cleanup, nil
}
