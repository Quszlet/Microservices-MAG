package repository

import (
	"time"

	"github.com/Quszlet/schedule_service/internal/models"
	"github.com/jmoiron/sqlx"
)

type Schedule interface {
	Create(ts models.TimeSlot) (int, error)
	Update(timeSlot map[string]any, where string) error
	GetTimeSlotById(timeSlotId int) (models.TimeSlot, error)
	GetScheduleByDoctorIdAndDate(doctorId int, date time.Time) ([]models.TimeSlot, error)
	GetScheduleByDoctorId(doctorId int) ([]models.TimeSlot, error)
	GetScheduleFields(timeSlotId int, fields []string) (map[string]any, error)
	ChangeDoctorActivity(doctorId int, active bool) error
	DeleteScheduleByDoctorId(doctorId int) error
	Delete(timeSlotId int) error
}

type Repository struct {
	Schedule
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Schedule: NewSchedulePostgres(db)}
}
