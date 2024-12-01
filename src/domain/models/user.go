package models

import (
	"time"
)

type User struct {
	ID        uint
  Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Groups    []Group `gorm:"many2many:user_groups"`
}
