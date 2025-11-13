package service

import (
	"github.com/Quszlet/doctors_service/internal/models"
	"github.com/Quszlet/doctors_service/internal/repository"
)


type Doctor interface {
	Create(d models.Doctor) (int, error)
	Update(d models.Doctor, where string) error
	Get(doctorId int) (models.Doctor, error)
	Delete(doctorId int) error
}

type Service struct {
	Doctor
}

func NewService(r *repository.Repository) *Service {
	return &Service{Doctor: NewDoctorService(r)}
}