package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
