package repository

import (
	"errors"
	"fmt"

	"github.com/Quszlet/libs/utils_sql"
	"github.com/Quszlet/doctors_service/internal/models"
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

	fmt.Println(q);
	
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
