package model

type Restaurant struct {
	RestaurantID string `gorm:"primaryKey" json:"restaurantID"`
	Password     string `gorm:"not null" json:"password"`
	FirstName    string `gorm:"not null" json:"firstName"`
	LastName     string `gorm:"not null" json:"lastName"`
	Designation  string `gorm:"not null" json:"designation"`
}
