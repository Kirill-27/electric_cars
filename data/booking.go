package data

import "time"

type Booking struct {
	Id              int       `json:"id" db:"id"`
	CustomerId      int       `json:"customer_id" db:"customer_ids" binding:"required"`
	StationId       string    `json:"station_id" db:"station_id" binding:"required"`
	PaymentMethodId string    `json:"payment_method_id" db:"payment_method_id" binding:"required"`
	StartTime       time.Time `json:"start_time" db:"start_time" binding:"required"`
	EndTime         time.Time `json:"end_time" db:"end_time" binding:"required"`
	IsPaid          bool      `json:"is_paid" db:"is_paid" binding:"required"`
	IsActive        bool      `json:"is_active" db:"is_active" binding:"required"`
}
