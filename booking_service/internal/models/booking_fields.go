package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BookingFields struct {
	TimeSlotID int      `json:"time_slot_id" db:"time_slot_id"`
	Fields     []string `json:"fields"`
}

func (bf BookingFields) Validate() error {
	return validation.ValidateStruct(&bf,
		validation.Field(&bf.TimeSlotID, validation.Required),
		validation.Field(&bf.Fields, validation.Required),
	)
}
