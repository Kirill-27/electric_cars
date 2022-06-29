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
	GetCustomerById(id int) (*data.Customer, error)
}

type Station interface {
	CreateStation(station data.Station) (int, error)
	GetAllStations() ([]data.Station, error)
	GetNearestStation(x, y float64) (*data.Station, error)
	GetStationById(stationId int) (*data.Station, error)
	DeleteStation(stationId int) error
	UpdateStation(station data.Station) error
}

type PaymentMethod interface {
	Create(paymentMethod data.PaymentMethod) (int, error)
	GetAll() ([]data.PaymentMethod, error)
	GetById(paymentMethodId int) (*data.PaymentMethod, error)
	Delete(paymentMethodId int) error
	Update(paymentMethod data.PaymentMethod) error
}

type Service struct {
	Authorization
	Station
	PaymentMethod
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Station:       NewStationService(repos),
		PaymentMethod: NewPaymentMethodService(repos),
	}
}
