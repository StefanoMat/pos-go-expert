package entity

import (
	"errors"
	"time"

	"github.com/stefanomat/pos-go-expert/9-API/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrIDIsRequired    = errors.New("ID is required")
	ErrInvalidID       = errors.New("ID is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := p.IsValid(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) IsValid() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}
