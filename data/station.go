package data

type Station struct {
	Id          int     `json:"id" db:"id"`
	IsFree      bool    `json:"is_free" db:"is_free" binding:"required"`
	LocationX   float64 `json:"location_x" db:"location_x" binding:"required"`
	LocationY   float64 `json:"location_y" db:"location_y" binding:"required"`
	ValuePerMin int     `json:"value_per_min" db:"value_per_min" binding:"required"`
	IsActive    bool    `json:"is_active" db:"is_active" binding:"required"`
}
