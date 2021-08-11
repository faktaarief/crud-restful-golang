package models

import "time"

type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" validate:"required" json:"title"`
	Content   string    `gorm:"type:text;not null" validate:"required" json:"content"`
	Author    User      `json:"author" validate:"-"`
	AuthorID  uint32    `gorm:"not null" validate:"required" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
