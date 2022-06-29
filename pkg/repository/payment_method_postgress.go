package repository

import (
	"database/sql"
	"fmt"
	"github.com/Kirill-27/electric_cars/data"
	"github.com/jmoiron/sqlx"
)

type PaymentMethodPostgres struct {
	db *sqlx.DB
}

func NewPaymentMethodPostgres(db *sqlx.DB) *PaymentMethodPostgres {
	return &PaymentMethodPostgres{db: db}
}

func (r *PaymentMethodPostgres) Create(paymentMethod data.PaymentMethod) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (customer_id, details, is_active) VALUES ($1, $2, $3) RETURNING id", paymentMethodsTable)

	row := r.db.QueryRow(query,
		paymentMethod.CustomerId,
		paymentMethod.Details,
		paymentMethod.IsActive)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PaymentMethodPostgres) GetAll() ([]data.PaymentMethod, error) {
	var paymentMethods []data.PaymentMethod

	query := fmt.Sprintf("SELECT * FROM %s", paymentMethodsTable)
	err := r.db.Select(&paymentMethods, query)

	return paymentMethods, err
}

func (r *PaymentMethodPostgres) GetById(paymentMethodId int) (*data.PaymentMethod, error) {
	var paymentMethod data.PaymentMethod

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", paymentMethodsTable)
	err := r.db.Get(&paymentMethod, query, paymentMethodId)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &paymentMethod, err
}

func (r *PaymentMethodPostgres) Delete(paymentMethodId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", paymentMethodsTable)
	_, err := r.db.Exec(query, paymentMethodId)
	return err
}

func (r *PaymentMethodPostgres) Update(paymentMethod data.PaymentMethod) error {
	query := fmt.Sprintf("UPDATE %s SET customer_id=$1, details=$2, is_active=$3 WHERE id=$4 ", paymentMethodsTable)
	_, err := r.db.Exec(query, paymentMethod.CustomerId, paymentMethod.Details, paymentMethod.IsActive, paymentMethod.Id)
	return err
}
