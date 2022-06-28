package data

type Customer struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"id" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Phone    string `json:"phone" db:"phone" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Role     int    `json:"role" db:"role" binding:"required"`
}
