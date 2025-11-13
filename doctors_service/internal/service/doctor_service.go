package service

import (
	"github.com/Quszlet/doctors_service/internal/models"
	"github.com/Quszlet/doctors_service/internal/repository"

	"github.com/fatih/structs"
)

type DoctorService struct {
	repo repository.Doctor
}

func NewDoctorService(r *repository.Repository) *DoctorService {
	return &DoctorService{repo: r.Doctor}
}

func (ds *DoctorService) Create(d models.Doctor) (int, error) {
	return ds.repo.Create(d)
}

func (ds *DoctorService) Update(d models.Doctor, where string) error {
	md := structs.Map(d)
	return ds.repo.Update(md, where)
}

func (ds *DoctorService) Get(doctorId int) (models.Doctor, error) {
	return ds.repo.Get(doctorId)
}

func (ds *DoctorService) Delete(doctorId int) error {
	return ds.repo.Delete(doctorId)
}
