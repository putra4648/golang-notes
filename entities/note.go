package entities

import "time"

type Note struct {
	ID      string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title   string
	Content string
	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time `gorm:"autoUpdateTime"`
}
