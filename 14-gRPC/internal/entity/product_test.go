package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	name := "iPhone"
	price := 500.0
	product, err := NewProduct(name, price)
	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, name, product.Name)
	require.Equal(t, price, product.Price)
}

func TestProductNameIsRequired(t *testing.T) {
	p, err := NewProduct("iPhone", 0.0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductPriceIsRequired(t *testing.T) {
	p, err := NewProduct("iPhone", 0.0)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("iPhone", 500.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.IsValid())
}
