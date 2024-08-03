package domains

import "time"

type User struct {
	Id        int    `gorm:"not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:100;not null;default:null"`
	Email     string `gorm:"unique;size:100;not null;uniqueIndex;default:null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
