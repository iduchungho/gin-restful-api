package model

type Payment struct {
	Id            string `gorm:"primaryKey" json:"id"`
	OrderID       string `gorm:"type:varchar(255);not null" json:"order_id"`
	PaymentType   string `gorm:"type:varchar(255);not null" json:"payment_type"`
	PaymentStatus string `gorm:"type:varchar(255);not null" json:"payment_status"`
	TimeStamp     string `gorm:"type:varchar(255);not null" json:"time_stamp"`
}
