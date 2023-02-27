package model

type PaymentDetails struct {
	PaymentID      string `gorm:"primaryKey" json:"paymentID"`
	CustomerID     string `gorm:"not null" json:"customerID"`
	CardNumber     string `gorm:"not null" json:"cardNumber"`
	CardHolderName string `gorm:"not null" json:"cardHolderName"`
	Cvv            string `gorm:"not null" json:"cvv"`
	ExpMonth       string `gorm:"not null" json:"expMonth"`
	ExpYear        string `gorm:"not null" json:"expYear"`
	TimeStamp      string `gorm:"not null" json:"timeStamp"`
}
