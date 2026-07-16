package models

import "time"

type Posts struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title" binding:"required,min=20"`
	Content     string    `gorm:"type:text;not null" json:"content" binding:"required,min=200"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category" binding:"required,min=3"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null" json:"status" binding:"required,oneof=Publish Draft Thrash"`
}
