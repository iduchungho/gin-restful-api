package model

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerID"`
	FirstName  string `gorm:"not null" json:"firstName"`
	LastName   string `gorm:"not null" json:"lastName"`
	EmailID    string `gorm:"not null" gorm:"unique" json:"emailID"`
	Password   string `gorm:"not null" json:"password"`
	PhoneNo    string `gorm:"not null" json:"phoneNo"`
	State      string `gorm:"not null" json:"state"`
	Landmark   string `gorm:"not null" json:"landmark"`
	PinCode    string `gorm:"not null" json:"pinCode"`
}
