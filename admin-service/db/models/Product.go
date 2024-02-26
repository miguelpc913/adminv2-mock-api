package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ProductId int            `json:"productId" gorm:"primary_key"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
