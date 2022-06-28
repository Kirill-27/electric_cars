package data

type PaymentMethod struct {
	Id         int    `json:"-" db:"id"`
	CustomerId int    `json:"customer_id" db:"customer_id" binding:"required"`
	Details    string `json:"details" db:"details" binding:"required"`
}
