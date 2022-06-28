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

func (s *StationService) Create(station data.Station) (int, error) {
	return s.repo.Create(station)
}

func (s *StationService) GetAll() ([]data.Station, error) {
	return s.repo.GetAll()
}

func (s *StationService) GetById(stationId int) (data.Station, error) {
	return s.repo.GetById(stationId)
}

func (s *StationService) Delete(stationId int) error {
	return s.repo.Delete(stationId)
}

func (s *StationService) Update(stationId, station data.Station) error {
	return s.repo.Update(stationId, station)
}
