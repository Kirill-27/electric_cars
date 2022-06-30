package repository

import (
	"database/sql"
	"fmt"
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
	"math"
)

type StationPostgres struct {
	db *sqlx.DB
}

const maxLength = float64(1000000000)

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

func (r *StationPostgres) UpdateStation(station data.Station) error {
	query := fmt.Sprintf("UPDATE %s SET is_free=$1, location_x=$2, location_y=$3, value_per_min=$4, is_active=$5 WHERE id=$6 ", stationsTable)
	_, err := r.db.Exec(query, station.IsFree, station.LocationX, station.LocationY, station.ValuePerMin, station.IsActive, station.Id)
	return err
}

func (r *StationPostgres) DeleteStation(stationId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", stationsTable)
	_, err := r.db.Exec(query, stationId)
	return err
}

func (r *StationPostgres) GetNearestStation(x, y float64) (*data.Station, error) {
	var stations []data.Station

	query := fmt.Sprintf("SELECT * FROM %s", stationsTable)
	err := r.db.Select(&stations, query)
	if err != nil {
		return nil, err
	}

	var res data.Station
	length := maxLength
	for _, st := range stations {
		dif := math.Sqrt((x-st.LocationX)*(x-st.LocationX) + (y-st.LocationY)*(y-st.LocationY))
		if dif < length {
			length = dif
			res = st
		}
	}
	return &res, nil
}
