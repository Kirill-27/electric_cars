package repository

import (
	"fmt"
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateCustomer(customer data.Customer) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, "+
		"password, phone) values ($1, $2, $3, $4, $5) RETURNING id", customersTable)

	row := r.db.QueryRow(query,
		customer.Name,
		customer.Username,
		customer.Email,
		customer.Password,
		customer.Phone)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetCustomer(username, password string) (data.Customer, error) {
	var customer data.Customer
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", customersTable)
	err := r.db.Get(&customer, query, username, password)

	return customer, err
}
