package data

type PaymentMethod struct {
	Id         int    `json:"id" db:"id"`
	CustomerId int    `json:"customer_id" db:"customer_id" binding:"required"`
	Details    string `json:"details" db:"details" binding:"required"`
	IsActive   bool   `json:"is_active" db:"is_active"`
}
