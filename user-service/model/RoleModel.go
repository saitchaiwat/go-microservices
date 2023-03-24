package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey;not null;autoIncrement" db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
