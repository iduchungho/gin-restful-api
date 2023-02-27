package model

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customer_id"`
	FirstName  string `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName   string `gorm:"type:varchar(255);not null" json:"lastName"`
	EmailID    string `gorm:"type:varchar(255);not null" gorm:"unique" json:"emailID"`
	Password   string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNo    string `gorm:"type:varchar(255);not null" json:"phoneNo"`
	State      string `gorm:"type:varchar(255);not null" json:"state"`
	Landmark   string `gorm:"type:varchar(255);not null" json:"landmark"`
	PinCode    string `gorm:"type:varchar(255);not null" json:"pinCode"`
}
