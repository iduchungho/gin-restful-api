package model

type Restaurant struct {
	RestaurantID string `gorm:"primaryKey" json:"restaurant_id"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	FirstName    string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName     string `gorm:"type:varchar(255);not null" json:"last_name"`
	Designation  string `gorm:"type:varchar(255);not null" json:"designation"`
}
