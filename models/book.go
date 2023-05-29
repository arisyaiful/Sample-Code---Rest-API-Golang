package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Id_category uint      `jdon:"id_category"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateBook struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Id_category uint   `json:"id_category" binding:"required"`
}

type UpdateBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
