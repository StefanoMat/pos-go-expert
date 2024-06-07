package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	SerialNumber SerialNumber
	CategoryID   int
	Category     Category
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
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      100.99,
	// 	CategoryID: category.ID,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "1234567",
	// 	ProductID: 1,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	var products []Product
	db.Preload("Category").Find(&products)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, product.Category.Name, "SerialNumber:", product.SerialNumber.Number)

		}
	}

}
