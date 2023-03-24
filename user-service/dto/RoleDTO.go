package dto

import "time"

type GetAllRole struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type GetRole struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateRole struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRole struct {
	Name string `json:"name" validate:"required"`
}
