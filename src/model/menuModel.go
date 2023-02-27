package model

type Menu struct {
	MenuID   string `gorm:"primaryKey" json:"menuID"`
	MenuName string `gorm:"not null" json:"menuName"`
	Price    string `gorm:"not null" json:"price"`
}
