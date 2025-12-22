package repository

import (
	"github.com/Quszlet/booking_service/internal/models"
	"github.com/jmoiron/sqlx"
)

type Booking interface {
	Create(rb models.RequestBooking) (int, error)
	GetBookingById(bookingId int) (models.RequestBooking, error)
	GetBookingFields(timeSlotId int, fields []string) (map[string]any, error)
	CancelledBookingByTimeSlotId(timeSlotId int) error
	DeleteBookingByTimeSlotId(bookingId int) error
}

type Repository struct {
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Booking: NewBookingPostgres(db)}
}
