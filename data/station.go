package data

type Station struct {
	Id          int     `json:"-" db:"id"`
	IsFree      bool    `json:"is_free" binding:"required"`
	LocationX   float32 `json:"location_x" binding:"required"`
	LocationY   float32 `json:"location_y" binding:"required"`
	ValuePerMin int     `json:"value_per_min" binding:"required"`
}
