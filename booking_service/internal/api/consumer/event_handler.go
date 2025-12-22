package consumer

import (
	"context"
	"log/slog"

	"github.com/Quszlet/booking_service/internal/service"
	kafkalib "github.com/Quszlet/libs/kafka"
)

type EventHandler struct {
	service *service.Service
}

func NewEventHandler(s *service.Service) *EventHandler {
	return &EventHandler{service: s}
}

func ConsumeKafkaEvents(ctx context.Context, kafkaConsumer *kafkalib.Consumer, eventHandler *EventHandler) error {
	return kafkaConsumer.StartConsuming(ctx, func(ctx context.Context, event *kafkalib.Event) error {
		switch event.EventType {
		case DoctorDeactiveteScheduleEventType:
			var payload DoctorDeactiveteScheduleData
			if err := event.DecodeData(&payload); err != nil {
				slog.Error("failed to decode doctor.active_changed event payload", "error", err.Error())
				return err
			}

			return eventHandler.HandleDoctorDeactiveSchedule(ctx, &DoctorDeactiveteScheduleEvent{
				EventType: event.EventType,
				Timestamp: event.Timestamp,
				Data:      payload,
			})
		case TimeSlotDeletedEventType:
			var payload TimeSlotDeletedData
			if err := event.DecodeData(&payload); err != nil {
				slog.Error("failed to decode doctor.deleted event payload", "error", err.Error())
				return err
			}

			return eventHandler.HandleDeletedSchedule(ctx, &TimeSlotDeletedEvent{
				EventType: event.EventType,
				Timestamp: event.Timestamp,
				Data:      payload,
			})
		default:
			return nil
		}
	})
}

func (h *EventHandler) HandleDoctorDeactiveSchedule(ctx context.Context, event *DoctorDeactiveteScheduleEvent) error {
	slog.Info("Handling schedule.deactivate event")

	for _, value := range event.Data.TimeSlotsID {
		err := h.service.CancelledBookingByTimeSlotId(value)
		if err != nil {
			slog.Error(err.Error())
		}
	}

	return nil
}

func (h *EventHandler) HandleDeletedSchedule(ctx context.Context, event *TimeSlotDeletedEvent) error {
	slog.Info("Handling schedule.deleted event")

	for _, value := range event.Data.TimeSlotsID {
		err := h.service.DeleteByBookingID(value)
		if err != nil {
			return err
		}
	}

	return nil
}
