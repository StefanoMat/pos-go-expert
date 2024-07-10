package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stefanomat/pos-go-expert/9-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindAllProduc(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		productDB.Create(product)
	}
	products, err := productDB.FindAll(1, 10, "asc")
	require.Nil(t, err)
	assert.Len(t, products, 10)
	require.Equal(t, "Product 1", products[0].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	db.Create(product)
	productDB := NewProduct(db)
	assert.NoError(t, err)
	productDB.Create(product)
	productFound, err := productDB.FindByID(product.ID.String())
	require.Nil(t, err)
	require.Equal(t, product.ID, productFound.ID)
	require.Equal(t, product.Name, productFound.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	db.Create(product)
	productDB := NewProduct(db)
	assert.NoError(t, err)
	productDB.Create(product)
	product.Name = "Product 2"
	productDB.Update(product)
	productFound, err := productDB.FindByID(product.ID.String())
	require.Nil(t, err)
	require.Equal(t, product.Name, productFound.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	db.Create(product)
	productDB := NewProduct(db)
	assert.NoError(t, err)
	productDB.Create(product)
	productDB.Delete(product.ID.String())
	_, err = productDB.FindByID(product.ID.String())
	require.Error(t, err)
}
