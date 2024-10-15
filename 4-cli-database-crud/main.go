package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uint `gorm:"primarykey;autoIncrement"`
	Name string
	Price float32
}

func main() {
	dsn := "root:password@tcp(0.0.0.0:3306)/public?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to get DB")
	}

	defer sqlDB.Close()

	db.AutoMigrate(&Product{})

	// db.Create(&Product{Name: "P01", Price: 10.2})

	var products []Product

	db.Select("id", "name", "price").Find(&products)

	for _, product := range products {
		log.Printf("{ID: %d, Name: '%s', Price: %.2f}", product.ID, product.Name, product.Price)
	}
}
