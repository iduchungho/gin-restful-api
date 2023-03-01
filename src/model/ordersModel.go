package model

type Orders struct {
	OrderID     string `gorm:"primaryKey" json:"order_id"`
	CustomerID  string `gorm:"type:varchar(255);not null" json:"customer_id"`
	MenuID      string `gorm:"type:varchar(255);not null" json:"menu_id"`
	Quantity    string `gorm:"type:varchar(255);not null" json:"quantity"`
	OrderStatus string `gorm:"type:varchar(255);not null" json:"order_status"`
	TimeStamp   string `gorm:"type:varchar(255);not null" json:"time_stamp"`
}
