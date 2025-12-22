package service

import (
	"time"

	"github.com/Quszlet/schedule_service/internal/models"
	"github.com/Quszlet/schedule_service/internal/repository"

	"github.com/fatih/structs"
)

type ScheduleService struct {
	repo repository.Schedule
}

func NewDoctorService(r *repository.Repository) *ScheduleService {
	return &ScheduleService{repo: r.Schedule}
}

func (ds *ScheduleService) Create(ts models.TimeSlot) (int, error) {
	return ds.repo.Create(ts)
}

func (ds *ScheduleService) CreateScheduleForDoctor(doctorId int, date time.Time, weeks int) error {
	const daySlots = 5
	dayIncrement := 1
	const countVisits = 11

	for dayIncrement <= weeks*daySlots {
		day := date.Weekday()
		if day != time.Saturday && day != time.Sunday {
			startTimeDay := time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)
			endTimeDay := startTimeDay.Add(time.Minute * 60)
			for i := 0; i < countVisits; i++ {
				startTimeDay := endTimeDay
				endTimeDay = endTimeDay.Add(time.Minute * 60)
				timeSlot := models.TimeSlot{
					DoctorID:    doctorId,
					Date:        date,
					StartTime:   startTimeDay,
					EndTime:     endTimeDay,
					IsAvailable: true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				if _, err := ds.repo.Create(timeSlot); err != nil {
					return err
				}
			}
		}

		date = date.AddDate(0, 0, 1)
		dayIncrement++
	}

	return nil
}

func (ds *ScheduleService) Update(ts models.TimeSlot, where string) error {
	md := structs.Map(ts)
	return ds.repo.Update(md, where)
}

func (ds *ScheduleService) GetTimeSlotById(timeSlotId int) (models.TimeSlot, error) {
	return ds.repo.GetTimeSlotById(timeSlotId)
}

func (ds *ScheduleService) GetScheduleByDoctorIdAndDate(doctorId int, date time.Time) ([]models.TimeSlot, error) {
	return ds.repo.GetScheduleByDoctorIdAndDate(doctorId, date)
}

func (ds *ScheduleService) GetScheduleByDoctorId(doctorId int) ([]models.TimeSlot, error) {
	return ds.repo.GetScheduleByDoctorId(doctorId)
}

func (ds *ScheduleService) GetTimeSlotFields(timeSlotId int, fields []string) (map[string]any, error) {
	return ds.repo.GetScheduleFields(timeSlotId, fields)
}

func (ds *ScheduleService) Delete(doctorId int) error {
	return ds.repo.Delete(doctorId)
}

func (ds *ScheduleService) ChangeDoctorActivity(doctorId int, active bool) error {
	return ds.repo.ChangeDoctorActivity(doctorId, active)
}

func (ds *ScheduleService) DeleteScheduleByDoctorId(doctorId int) error {
	return ds.repo.DeleteScheduleByDoctorId(doctorId)
}
