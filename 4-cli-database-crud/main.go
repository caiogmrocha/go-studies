package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	godotenv.Load()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DEFAULT"),
	)
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
