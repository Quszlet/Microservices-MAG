package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RequestSchedule struct {
	DoctorID int    `json:"doctor_id"`
	Date     time.Time `json:"date`
}

func (rs RequestSchedule) ValidateCreate() error {
	return validation.ValidateStruct(&rs,
		validation.Field(&rs.DoctorID, validation.Required),
		validation.Field(&rs.Date, validation.Required),
	)
}
