package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Doctor struct {
	Id           uint   `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Surname      string `json:"surname" db:"surname"`
	Patronymic   string `json:"patronymic" db:"patronymic"`
	Specialty    string `json:"specialty" db:"specialty"`
	Clinic       string `json:"clinic" db:"clinic"`
	Status *bool  `json:"status_active" db:"status"`
}

// Обычная валидация для Create (все обязательны)
func (d Doctor) ValidateCreate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Surname, validation.Required),
		validation.Field(&d.Specialty, validation.Required),
		validation.Field(&d.Clinic, validation.Required),
	)
}

// PATCH: валидируем только присланные поля
func (d Doctor) ValidatePatch(provided map[string]bool) error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Id, validation.Required),
		validation.Field(&d.Name,
			validation.When(provided["name"],
				validation.Required, validation.RuneLength(1, 50),
			).Else(validation.Skip),
		),
		validation.Field(&d.Surname,
			validation.When(provided["surname"],
				validation.Required, validation.RuneLength(1, 50),
			).Else(validation.Skip),
		),
		validation.Field(&d.Patronymic,
			validation.When(provided["patronymic"],
				validation.RuneLength(0, 50),
			).Else(validation.Skip),
		),
		validation.Field(&d.Specialty,
			validation.When(provided["specialty"],
				validation.Required, validation.RuneLength(1, 50),
			).Else(validation.Skip),
		),
		validation.Field(&d.Clinic,
			validation.When(provided["clinic"],
				validation.Required, validation.RuneLength(1, 50),
			).Else(validation.Skip),
		),
	)
}
