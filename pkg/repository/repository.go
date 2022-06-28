package repository

import (
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateCustomer(customer data.Customer) (int, error)
	GetCustomer(username, password string) (data.Customer, error)
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

type Repository struct {
	Authorization
	TodoList
	TodoItem
	Station
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      new(TodoList),
		TodoItem:      new(TodoItem),
		Station:       NewStationPostgres(db),
	}
}
