package repository

import (
	"database/sql"
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

func (r *StationPostgres) CreateStation(station data.Station) (int, error) {
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

func (r *StationPostgres) GetAllStations() ([]data.Station, error) {
	var stations []data.Station

	query := fmt.Sprintf("SELECT * FROM %s", stationsTable)
	err := r.db.Select(&stations, query)

	return stations, err
}
func (r *StationPostgres) GetStationById(stationId int) (*data.Station, error) {
	var station data.Station

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", stationsTable)
	err := r.db.Get(&station, query, stationId)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &station, err
}

//TODO chech this
func (r *StationPostgres) UpdateStation(station data.Station) error {
	query := fmt.Sprintf("UPDATE %s SET is_free=$1, location_x=$2, location_y=$3, value_per_min=$4 WHERE id=$5 ", stationsTable)
	_, err := r.db.Exec(query, station.IsFree, station.LocationX, station.LocationY, station.ValuePerMin, station.Id)
	return err
}

func (r *StationPostgres) DeleteStation(stationId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", stationsTable)
	_, err := r.db.Exec(query, stationId)
	return err
}
