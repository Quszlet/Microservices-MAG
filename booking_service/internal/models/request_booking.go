package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RequestBooking struct {
	TimeSlotID int `json:"time_slot_id"`
	PatientID  int `json:"patient_id"`
}

func (rs RequestBooking) ValidateCreate() error {
	return validation.ValidateStruct(&rs,
		validation.Field(&rs.TimeSlotID, validation.Required),
		validation.Field(&rs.PatientID, validation.Required),
	)
}
