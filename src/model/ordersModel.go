package model

type Orders struct {
	OrderID     string `gorm:"primaryKey" json:"orderID"`
	CustomerID  string `gorm:"not null" json:"customerID"`
	MenuID      string `gorm:"not null" json:"menuID"`
	Quantity    string `gorm:"not null" json:"quantity"`
	OrderStatus string `gorm:"not null" json:"orderStatus"`
	TimeStamp   string `gorm:"not null" json:"timeStamp"`
}
