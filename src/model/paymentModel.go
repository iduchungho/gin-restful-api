package model

type Payment struct {
	Id            string `gorm:"primaryKey" json:"id"`
	OrderID       string `gorm:"not null" json:"orderID"`
	PaymentType   string `gorm:"not null" json:"paymentType"`
	PaymentStatus string `gorm:"not null" json:"paymentStatus"`
	TimeStamp     string `gorm:"not null" json:"timeStamp"`
}
