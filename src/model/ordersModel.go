package model

type Orders struct {
	OrderID string `gorm:"primaryKey" json:"orderID"`
}
