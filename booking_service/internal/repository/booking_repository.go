package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Quszlet/booking_service/internal/models"
	utilssql "github.com/Quszlet/libs/utils_sql"
	"github.com/jmoiron/sqlx"
)

type BookingPostgres struct {
	db *sqlx.DB
}

func NewBookingPostgres(db *sqlx.DB) *BookingPostgres {
	return &BookingPostgres{db: db}
}

func (up *BookingPostgres) Create(rb models.RequestBooking) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (patient_id, time_slot_id, cancelled,  
	created_at) values ($1, $2, $3, $4) RETURNING id`, BookingTable)

	row := up.db.QueryRow(query,
		rb.PatientID,
		rb.TimeSlotID,
		false,
		time.Now(),
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (up *BookingPostgres) GetBookingById(bookingId int) (models.RequestBooking, error) {
	var booking models.RequestBooking
	query := fmt.Sprintf("SELECT * FROM %s WHERE time_slot_id = $1", BookingTable)
	err := up.db.Get(&booking, query, bookingId)
	return booking, err
}

func (up *BookingPostgres) GetBookingFields(timeSlotId int, fields []string) (map[string]any, error) {
	result := make(map[string]any)
	where := fmt.Sprintf("time_slot_id = %d", timeSlotId)
	q, args := utilssql.BuildGetQuery(BookingTable, fields, where)

	res, err := up.db.NamedQuery(q, args)
	if err != nil {
		return map[string]any{}, err
	}

	defer res.Close()

	if res.Next() {
		if err = res.MapScan(result); err != nil {
			return map[string]any{}, err
		}
	} else {
		return map[string]any{}, errors.New("booking with this timeSlotId does not exist")
	}

	return result, nil
}

func (up *BookingPostgres) CancelledBookingByTimeSlotId(timeSlotId int) error {
	query := fmt.Sprintf("UPDATE %s SET cancelled=true WHERE time_slot_id = $1", BookingTable)
	res, err := up.db.Exec(query, timeSlotId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("booking with this ID does not exist")
	}

	return err
}

func (up *BookingPostgres) DeleteBookingByTimeSlotId(bookingId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE time_slot_id = $1", BookingTable)
	res, err := up.db.Exec(query, bookingId)
	if err != nil {
		return err
	}

	affRows, err := res.RowsAffected()
	if affRows == 0 {
		return errors.New("booking with this ID does not exist")
	}

	return err
}
