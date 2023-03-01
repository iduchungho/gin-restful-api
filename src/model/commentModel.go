package model

type Comment struct {
	CommentID string `gorm:"primaryKey" json:"comment_id"`
	UserID    string `gorm:"type:varchar(255);not null" json:"user_id"`
	Content   string `gorm:"type:varchar(255);not null" json:"content"`
	Rating    string `gorm:"type:varchar(255);not null" json:"rating"`
}
