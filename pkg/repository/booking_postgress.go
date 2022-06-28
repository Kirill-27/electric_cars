package repository

import (
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

func (r *BookingPostgres) Create(station data.Station) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (is_free, location_x, location_y, value_per_min) "+
		"VALUES ($1, $2, $3, $4) RETURNING id", stationsTable)

	row := r.db.QueryRow(query,
		station.IsFree,
		station.LocationX,
		station.LocationY,
		station.ValuePerMin)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BookingPostgres) GetAll() ([]data.Station, error) {
	var stations []data.Station

	query := fmt.Sprintf("SELECT * FROM %s", stationsTable)
	err := r.db.Select(&stations, query)

	return stations, err
}
func (r *BookingPostgres) GetById(stationId int) (data.Station, error) {
	var station data.Station

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", stationsTable)
	err := r.db.Get(&station, query, stationId)

	return station, err
}
func (r *BookingPostgres) Delete(stationId int) error {
	return nil
}
func (r *BookingPostgres) Update(stationId, station data.Station) error {
	return nil
}
