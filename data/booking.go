package data

type Booking struct {
	Id              int    `json:"-" db:"id"`
	UserId          int    `json:"user_id" binding:"required"`
	StationId       string `json:"station_id" binding:"required"`
	StartTime       bool   `json:"start_time" binding:"required"`
	EndTime         string `json:"end_time" binding:"required"`
	PaymentMethodId string `json:"payment_method_id" binding:"required"`
	IsPaid          bool   `json:"is_paid" binding:"required"`
}
