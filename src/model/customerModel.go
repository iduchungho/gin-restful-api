package model

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customer_id"`
	FirstName  string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName   string `gorm:"type:varchar(255);not null" json:"last_name"`
	EmailID    string `gorm:"type:varchar(255);not null;unique" json:"email_id"`
	Password   string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNo    string `gorm:"type:varchar(255);not null" json:"phone_no"`
	State      string `gorm:"type:varchar(255);not null" json:"state"`
	Landmark   string `gorm:"type:varchar(255);not null" json:"landmark"`
	Role       string `gorm:"type:varchar(255);not null" json:"role"`
}
