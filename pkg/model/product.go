package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null;"`
	Provider    string  `json:"provider" gorm:"not null;"`
	Quantity    uint16  `json:"quantity" gorm:"not null;"`
	Price       float32 `json:"price" gorm:"not null;"`
	Description string  `json:"description" gorm:"not null;"`
}

type ProductService interface {
	Product(id uint) (*Product, error)
	Products() (*[]Product, error)
	CreateProduct(p *Product) (*Product, error)
	UpdateProduct(id uint, p *Product) (*Product, error)
	DeleteProduct(id uint) error
}
