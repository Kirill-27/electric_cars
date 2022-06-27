package data

type PaymentMethod struct {
	Id         int    `json:"-" db:"id"`
	CustomerId int    `json:"customer_id" binding:"required"`
	Details    string `json:"details" binding:"required"`
}
