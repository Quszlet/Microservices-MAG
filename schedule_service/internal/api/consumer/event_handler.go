package consumer

import (
	"context"
	"log/slog"
	"time"

	kafkalib "github.com/Quszlet/libs/kafka"
	"github.com/Quszlet/schedule_service/internal/kafka"
	"github.com/Quszlet/schedule_service/internal/service"
)

// EventHandler handles Kafka events for schedule service.
type EventHandler struct {
	service   *service.Service
	producers *kafka.Producers
}

func NewEventHandler(s *service.Service, p *kafka.Producers) *EventHandler {
	return &EventHandler{service: s, producers: p}
}

func ConsumeKafkaEvents(ctx context.Context, kafkaConsumer *kafkalib.Consumer, eventHandler *EventHandler) error {
	return kafkaConsumer.StartConsuming(ctx, func(ctx context.Context, event *kafkalib.Event) error {
		switch event.EventType {
		case DoctorCreatedEventType:
			var payload DoctorCreatedData
			if err := event.DecodeData(&payload); err != nil {
				slog.Error("failed to decode doctor.created event payload", "error", err.Error())
				return err
			}

			return eventHandler.HandleDoctorCreated(ctx, &DoctorCreatedEvent{
				EventType: event.EventType,
				Timestamp: event.Timestamp,
				Data:      payload,
			})
		case DoctorActiveChangedEventType:
			var payload DoctorActiveChangedData
			if err := event.DecodeData(&payload); err != nil {
				slog.Error("failed to decode doctor.active_changed event payload", "error", err.Error())
				return err
			}

			return eventHandler.HandleDoctorActiveChanged(ctx, &DoctorActiveChangedEvent{
				EventType: event.EventType,
				Timestamp: event.Timestamp,
				Data:      payload,
			})
		case DoctorDeletedEventType:
			var payload DoctorDeletedData
			if err := event.DecodeData(&payload); err != nil {
				slog.Error("failed to decode doctor.deleted event payload", "error", err.Error())
				return err
			}

			return eventHandler.HandleDoctorDeleted(ctx, &DoctorDeletedEvent{
				EventType: event.EventType,
				Timestamp: event.Timestamp,
				Data:      payload,
			})
		default:
			return nil
		}
	})
}

func (h *EventHandler) HandleDoctorCreated(ctx context.Context, event *DoctorCreatedEvent) error {
	slog.Info("Handling doctor.created event", "doctor_id", event.Data.DoctorID)

	err := h.service.CreateScheduleForDoctor(event.Data.DoctorID, time.Now(), 1)
	if err != nil {
		return err
	}

	return nil
}

func (h *EventHandler) HandleDoctorActiveChanged(ctx context.Context, event *DoctorActiveChangedEvent) error {
	slog.Info("Handling doctor.active_changed event", "doctor_id", event.Data.DoctorID)
	slog.Info("Handling doctor.active_changed event", "active", event.Data.Active)

	err := h.service.ChangeDoctorActivity(event.Data.DoctorID, event.Data.Active)
	if err != nil {
		return err
	}

	timeSlots, err := h.service.GetScheduleByDoctorId(event.Data.DoctorID)
	if err != nil {
		return err
	}

	var timeSlotsId []int
	for _, value := range timeSlots {
		timeSlotsId = append(timeSlotsId, value.ID)
	}

	h.producers.ScheduleEvents.SendMessage(ctx, "schedule.deactivate", map[string]any{
		"time_slots_id": timeSlotsId,
	})

	return nil
}

func (h *EventHandler) HandleDoctorDeleted(ctx context.Context, event *DoctorDeletedEvent) error {
	slog.Info("Handling doctor.deleted event", "doctor_id", event.Data.DoctorID)

	timeSlots, err := h.service.GetScheduleByDoctorId(event.Data.DoctorID)
	if err != nil {
		return err
	}

	err = h.service.DeleteScheduleByDoctorId(event.Data.DoctorID)
	if err != nil {
		return err
	}

	var timeSlotsId []int
	for _, value := range timeSlots {
		timeSlotsId = append(timeSlotsId, value.ID)
	}

	h.producers.ScheduleEvents.SendMessage(ctx, "schedule.time_slots_deleted", map[string]any{
		"time_slots_id": timeSlotsId,
	})

	return nil
}
