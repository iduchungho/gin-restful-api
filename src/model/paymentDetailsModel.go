package model

type PaymentDetails struct {
	PaymentID      string `gorm:"primaryKey" json:"payment_id"`
	CustomerID     string `gorm:"type:varchar(255);not null" json:"customer_id"`
	CardNumber     string `gorm:"type:varchar(255);not null" json:"card_number"`
	CardHolderName string `gorm:"type:varchar(255);not null" json:"card_holder_name"`
	Cvv            string `gorm:"type:varchar(255);not null" json:"cvv"`
	ExpMonth       string `gorm:"type:varchar(255);not null" json:"exp_month"`
	ExpYear        string `gorm:"type:varchar(255);not null" json:"exp_year"`
	TimeStamp      string `gorm:"type:varchar(255);not null" json:"time_stamp"`
}
