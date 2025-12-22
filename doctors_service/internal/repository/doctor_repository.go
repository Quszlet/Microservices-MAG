package repository

import (
	"errors"
	"fmt"

	"github.com/Quszlet/doctors_service/internal/models"
	utilssql "github.com/Quszlet/libs/utils_sql"
	"github.com/jmoiron/sqlx"
)

type DoctorPostgres struct {
	db *sqlx.DB
}

func NewDoctorPostgres(db *sqlx.DB) *DoctorPostgres {
	return &DoctorPostgres{db: db}
}

func (up *DoctorPostgres) Create(doctor models.Doctor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, specialty, clinic, status) values ($1, $2, $3, $4, $5, $6) RETURNING id", doctorsTable)

	row := up.db.QueryRow(query, doctor.Name, doctor.Surname, doctor.Patronymic, doctor.Specialty, doctor.Clinic, doctor.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (up *DoctorPostgres) Update(doctor map[string]any, where string) error {
	q, args := utilssql.BuildUpdateQuery(doctorsTable, doctor, where)

	_, err := up.db.NamedExec(q, args)
	if err != nil {
		return err
	}

	return nil
}

func (up *DoctorPostgres) Get(doctorId int) (models.Doctor, error) {
	var doctor models.Doctor
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", doctorsTable)
	err := up.db.Get(&doctor, query, doctorId)
	return doctor, err
}

func (up *DoctorPostgres) GetDoctorFields(doctorId int, fields []string) (map[string]any, error) {
	result := make(map[string]any)
	where := fmt.Sprintf("id = %d", doctorId)
	q, args := utilssql.BuildGetQuery(doctorsTable, fields, where)

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
		return map[string]any{}, errors.New("doctor with this ID does not exist")
	}

	return result, nil
}

func (up *DoctorPostgres) Delete(doctorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", doctorsTable)
	res, err := up.db.Exec(query, doctorId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("doctor with this ID does not exist")
	}

	return err
}
