package repository

import (
	"database/sql"
	"fmt"
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type BookingPostgres struct {
	db *sqlx.DB
}

func NewBookingPostgres(db *sqlx.DB) *BookingPostgres {
	return &BookingPostgres{db: db}
}

func (r *BookingPostgres) CreateBooking(booking data.Booking) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (customer_id, payment_method_id, station_id, is_paid, star_time, end_time) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", bookingsTable)

	row := r.db.QueryRow(query,
		booking.CustomerId,
		booking.PaymentMethodId,
		booking.StationId,
		booking.IsPaid,
		booking.StartTime,
		booking.EndTime)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BookingPostgres) GetAllBookings() ([]data.Booking, error) {
	var stations []data.Booking

	query := fmt.Sprintf("SELECT * FROM %s", bookingsTable)
	err := r.db.Select(&stations, query)

	return stations, err
}

func (r *BookingPostgres) GetBookingById(bookingId int) (*data.Booking, error) {
	var station data.Booking

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", bookingsTable)
	err := r.db.Get(&station, query, bookingId)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &station, err
}

func (r *BookingPostgres) DeleteBooking(bookingId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", bookingsTable)
	_, err := r.db.Exec(query, bookingId)
	return err
}

func (r *BookingPostgres) UpdateBooking(booking data.Booking) error {
	return nil
}
