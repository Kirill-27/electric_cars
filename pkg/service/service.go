package service

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/Kirill-27/electric_cars/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateCustomer(user data.Customer) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Station interface {
	Create(station data.Station) (int, error)
	GetAll() ([]data.Station, error)
	GetById(stationId int) (data.Station, error)
	Delete(stationId int) error
	Update(stationId, station data.Station) error
}

type Service struct {
	Authorization
	TodoList
	Station
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		TodoList:      new(TodoList),
		TodoItem:      new(TodoItem),
		Station:       NewStationService(repos),
	}
}
