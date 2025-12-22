package repository

import (
	"errors"
	"fmt"
	"time"

	utilssql "github.com/Quszlet/libs/utils_sql"
	"github.com/Quszlet/schedule_service/internal/models"
	"github.com/jmoiron/sqlx"
)

type SchedulePostgres struct {
	db *sqlx.DB
}

func NewSchedulePostgres(db *sqlx.DB) *SchedulePostgres {
	return &SchedulePostgres{db: db}
}

func (up *SchedulePostgres) Create(timeSlot models.TimeSlot) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (doctor_id, date, start_time, end_time, is_available, 
	created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, ScheduleTable)

	row := up.db.QueryRow(query,
		timeSlot.DoctorID,
		timeSlot.Date,
		timeSlot.StartTime,
		timeSlot.EndTime,
		timeSlot.IsAvailable,
		timeSlot.CreatedAt,
		timeSlot.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (up *SchedulePostgres) Update(timeSlot map[string]any, where string) error {
	q, args := utilssql.BuildUpdateQuery(ScheduleTable, timeSlot, where)

	_, err := up.db.NamedExec(q, args)
	if err != nil {
		return err
	}

	return nil
}

func (up *SchedulePostgres) GetTimeSlotById(timeSlotId int) (models.TimeSlot, error) {
	var timeSlot models.TimeSlot
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", ScheduleTable)
	err := up.db.Get(&timeSlot, query, timeSlotId)
	return timeSlot, err
}

func (up *SchedulePostgres) GetScheduleByDoctorIdAndDate(doctorId int, date time.Time) ([]models.TimeSlot, error) {
	var timeSlots []models.TimeSlot
	query := fmt.Sprintf("SELECT * FROM %s WHERE doctor_id = $1 AND date = $2", ScheduleTable)
	err := up.db.Select(&timeSlots, query, doctorId, date)
	return timeSlots, err
}

func (up *SchedulePostgres) GetScheduleByDoctorId(doctorId int) ([]models.TimeSlot, error) {
	var timeSlots []models.TimeSlot
	query := fmt.Sprintf("SELECT * FROM %s WHERE doctor_id = $1", ScheduleTable)
	err := up.db.Select(&timeSlots, query, doctorId)
	return timeSlots, err
}

func (up *SchedulePostgres) GetScheduleFields(timeSlotId int, fields []string) (map[string]any, error) {
	result := make(map[string]any)
	where := fmt.Sprintf("id = %d", timeSlotId)
	q, args := utilssql.BuildGetQuery(ScheduleTable, fields, where)

	res, err := up.db.NamedQuery(q, args)
	if err != nil {
		return map[string]any{}, err
	}

	defer res.Close()

	if res.Next() {
		err = res.MapScan(result)
		if err != nil {
			return map[string]any{}, err
		}
	} else {
		return map[string]any{}, errors.New("timeSlot with this ID does not exist")
	}

	return result, nil
}

func (up *SchedulePostgres) ChangeDoctorActivity(doctorId int, active bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_available=$1 WHERE doctor_id = $2", ScheduleTable)

	res, err := up.db.Exec(query, active, doctorId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("timeSlots with this doctor_id does not exist")
	}

	return err
}

func (up *SchedulePostgres) Delete(timeSlotId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ScheduleTable)
	res, err := up.db.Exec(query, timeSlotId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("timeSlot with this ID does not exist")
	}

	return err
}

func (up *SchedulePostgres) DeleteScheduleByDoctorId(doctorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE doctor_id = $1", ScheduleTable)
	res, err := up.db.Exec(query, doctorId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("timeSlots with this doctor_id does not exist")
	}

	return err
}
