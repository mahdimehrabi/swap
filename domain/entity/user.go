package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `gorm:"not null"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	CoinUsers []CoinUser `gorm:"foreignKey:UserID"`
}
