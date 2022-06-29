package repository

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateCustomer(customer data.Customer) (int, error)
	GetCustomer(username, password string) (*data.Customer, error)
	GetCustomerById(id int) (*data.Customer, error)
	UpdateCustomer(customer data.Customer) error
}

type Station interface {
	CreateStation(station data.Station) (int, error)
	GetAllStations() ([]data.Station, error)
	GetStationById(stationId int) (*data.Station, error)
	GetNearestStation(x, y float64) (*data.Station, error)
	DeleteStation(stationId int) error
	UpdateStation(station data.Station) error
}

type PaymentMethod interface {
	Create(paymentMethod data.PaymentMethod) (int, error)
	GetAll() ([]data.PaymentMethod, error)
	GetById(PaymentMethodId int) (*data.PaymentMethod, error)
	Delete(PaymentMethodId int) error
	Update(paymentMethod data.PaymentMethod) error
}

type Repository struct {
	PaymentMethod
	Authorization
	Station
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Station:       NewStationPostgres(db),
		PaymentMethod: NewPaymentMethodPostgres(db),
	}
}
