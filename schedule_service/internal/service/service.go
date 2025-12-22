package service

import (
	"time"

	"github.com/Quszlet/schedule_service/internal/models"
	"github.com/Quszlet/schedule_service/internal/repository"
)

type Schedule interface {
	Create(ts models.TimeSlot) (int, error)
	CreateScheduleForDoctor(doctorId int, date time.Time, weeks int) error
	Update(timeSlot models.TimeSlot, where string) error
	GetTimeSlotById(timeSlotId int) (models.TimeSlot, error)
	GetScheduleByDoctorIdAndDate(doctorId int, date time.Time) ([]models.TimeSlot, error)
	GetScheduleByDoctorId(doctorId int) ([]models.TimeSlot, error)
	GetTimeSlotFields(timeSlotId int, fields []string) (map[string]any, error)
	ChangeDoctorActivity(doctorId int, active bool) error
	DeleteScheduleByDoctorId(doctorId int) error
	Delete(timeSlotId int) error
}

type Service struct {
	Schedule
}

func NewService(r *repository.Repository) *Service {
	return &Service{Schedule: NewDoctorService(r)}
}
