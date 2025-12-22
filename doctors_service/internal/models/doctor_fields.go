package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DoctorFields struct {
	Id     int      `json:"id" db:"id"`
	Fields []string `json:"fields"`
}

func (df DoctorFields) Validate() error {
	return validation.ValidateStruct(&df,
		validation.Field(&df.Id, validation.Required),
		validation.Field(&df.Fields, validation.Required),
	)
}
