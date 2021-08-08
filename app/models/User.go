package models

type User struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string `gorm:"size:255;not null;unique" validate:"required" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" validate:"required" json:"email"`
	Password  string `gorm:"size:100;not null" validate:"required" json:"password"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
