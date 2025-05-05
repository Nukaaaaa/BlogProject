package models

type Post struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	UserID     uint     `json:"user_id"`
	User       User     `json:"-" gorm:"foreignKey:UserID"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"-" gorm:"foreignKey:CategoryID"`
}
