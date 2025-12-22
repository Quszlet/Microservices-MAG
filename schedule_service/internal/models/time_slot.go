package models

import "time"

type TimeSlot struct {
	ID          int     `json:"id" db:"id"`
	DoctorID    int     `json:"doctor_id" db:"doctor_id"`
	Date        time.Time `json:"date" db:"date"`
	StartTime   time.Time `json:"start_time" db:"start_time"`
	EndTime     time.Time `json:"end_time" db:"end_time"`
	IsAvailable bool      `json:"is_available" db:"is_available"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
