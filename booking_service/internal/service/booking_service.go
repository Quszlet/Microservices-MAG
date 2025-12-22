package service

import (
	"github.com/Quszlet/booking_service/internal/models"
	"github.com/Quszlet/booking_service/internal/repository"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(r *repository.Repository) *BookingService {
	return &BookingService{repo: r.Booking}
}

func (ds *BookingService) Create(rb models.RequestBooking) (int, error) {
	return ds.repo.Create(rb)
}

func (ds *BookingService) GetBookingById(bookingId int) (models.RequestBooking, error) {
	return ds.repo.GetBookingById(bookingId)
}

func (ds *BookingService) GetBookingFields(timeSlotId int, fields []string) (map[string]any, error) {
	return ds.repo.GetBookingFields(timeSlotId, fields)
}

func (ds *BookingService) CancelledBookingByTimeSlotId(timeSlotId int) error {
	return ds.repo.CancelledBookingByTimeSlotId(timeSlotId)
}

func (ds *BookingService) DeleteByBookingID(doctorId int) error {
	return ds.repo.DeleteBookingByTimeSlotId(doctorId)
}
