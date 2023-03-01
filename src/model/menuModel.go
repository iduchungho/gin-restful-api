package model

type Menu struct {
	MenuID   string `gorm:"primaryKey" json:"menu_id"`
	MenuName string `gorm:"type:varchar(255);not null" json:"menu_name"`
	Price    string `gorm:"type:varchar(255);not null" json:"price"`
}
