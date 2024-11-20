package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
  Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Groups    []Group `gorm:"many2many:user_groups"`
}

func (u User) getAllGroups(db *gorm.DB) []Group {
	var groups []Group

	db.Model(&u).Association("Groups").Find(&groups)

	return groups
}
