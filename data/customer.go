package data

const (
	CustomerRoleUser = iota + 1
	CustomerRoleAdmin
)

type Customer struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Phone    string `json:"phone" db:"phone" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Role     int    `json:"role" db:"role" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active"`
}
