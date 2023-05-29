package models

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategory struct {
	Name string `json:"name"`
}
