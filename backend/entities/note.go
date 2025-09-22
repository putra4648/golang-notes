package entities

import "time"

type Note struct {
	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDone    bool
}
