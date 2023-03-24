package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey;not null;autoIncrement" db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	RoleId    uint      `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
