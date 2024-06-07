package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Electronics"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Laptop",
	// 	Price:      1899.99,
	// 	CategoryID: category.ID,
	// })

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

}
