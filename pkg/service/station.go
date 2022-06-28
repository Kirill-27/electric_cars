package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

type StationService struct {
	repo repository.Station
}

func NewStationService(repo repository.Station) *StationService {
	return &StationService{repo: repo}
}

func (s *StationService) CreateStation(station data.Station) (int, error) {
	return s.repo.CreateStation(station)
}

func (s *StationService) GetAllStations() ([]data.Station, error) {
	return s.repo.GetAllStations()
}

func (s *StationService) GetStationById(stationId int) (*data.Station, error) {
	return s.repo.GetStationById(stationId)
}

func (s *StationService) DeleteStation(stationId int) error {
	return s.repo.DeleteStation(stationId)
}

func (s *StationService) UpdateStation(station data.Station) error {
	return s.repo.UpdateStation(station)
}
