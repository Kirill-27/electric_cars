package repository

import (
	"fmt"
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type StationPostgres struct {
	db *sqlx.DB
}

func NewStationPostgres(db *sqlx.DB) *StationPostgres {
	return &StationPostgres{db: db}
}

func (r *StationPostgres) Create(station data.Station) (int, error) {
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

func (r *StationPostgres) GetAll() ([]data.Station, error) {
	return []data.Station{}, nil
}
func (r *StationPostgres) GetById(stationId int) (data.Station, error) {
	var station data.Station

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", stationsTable)
	err := r.db.Get(&station, query, stationId)

	return station, err
}
func (r *StationPostgres) Delete(stationId int) error {
	return nil
}
func (r *StationPostgres) Update(stationId, station data.Station) error {
	return nil
}
