package repository

import (
	"github.com/Quszlet/doctors_service/internal/models"
	"github.com/jmoiron/sqlx"
)

type Doctor interface {
	Create(d models.Doctor) (int, error)
	Update(doctor map[string]any, where string) error
	Get(doctorId int) (models.Doctor, error)
	GetDoctorFields(doctorId int, fields []string) (map[string]any, error)
	Delete(doctorId int) error
}

type Repository struct {
	Doctor
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Doctor: NewDoctorPostgres(db)}
}