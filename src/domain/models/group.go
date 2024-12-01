package models

import "time"

type Group struct {
	ID        uint
  Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User `gorm:"many2many:user_groups"`
}
