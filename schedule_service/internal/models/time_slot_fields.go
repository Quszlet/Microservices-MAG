package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type TimeSlotFields struct {
	Id     int      `json:"id" db:"id"`
	Fields []string `json:"fields"`
}

func (tf TimeSlotFields) Validate() error {
	return validation.ValidateStruct(&tf,
		validation.Field(&tf.Id, validation.Required),
		validation.Field(&tf.Fields, validation.Required),
	)
}
