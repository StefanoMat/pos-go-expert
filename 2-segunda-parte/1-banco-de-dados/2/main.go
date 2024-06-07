package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{Name: "MacBook Pro", Price: 2199.99})

	// var products []Product
	// db.Find(&products)
	// fmt.Println(products)
	// db.Limit(2).Offset(1).Find(&products)
	// fmt.Println(products)

	// db.Where("price > ?", 2000).Find(&products)
	// db.Where("name LIKE ?", "%phone%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 2)
	// p.Name = "iPhone 14"
	// db.Save(&p)

	var p2 Product
	db.First(&p2, 7)
	fmt.Println(p2)
	db.Delete(&p2)

}
