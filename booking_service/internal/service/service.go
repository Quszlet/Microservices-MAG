package service

import (
	"github.com/Quszlet/booking_service/internal/models"
	"github.com/Quszlet/booking_service/internal/repository"
)

type Booking interface {
	Create(rb models.RequestBooking) (int, error)
	GetBookingById(bookingId int) (models.RequestBooking, error)
	GetBookingFields(timeSlotId int, fields []string) (map[string]any, error)
	CancelledBookingByTimeSlotId(timeSlotId int) error
	DeleteByBookingID(bookingId int) error
}

type Service struct {
	Booking
}

func NewService(r *repository.Repository) *Service {
	return &Service{Booking: NewBookingService(r)}
}
