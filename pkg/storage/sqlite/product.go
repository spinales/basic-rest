package sqlite

import (
	"basic-rest/pkg/model"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (s *ProductService) Product(id uint) (*model.Product, error) {
	var prod model.Product
	s.DB.First(&prod, id)
	return &prod, nil
}

func (s *ProductService) Products() (*[]model.Product, error) {
	var prods []model.Product
	s.DB.Find(&prods)
	return &prods, nil
}

func (s *ProductService) CreateProduct(p *model.Product) (*model.Product, error) {
	s.DB.Create(p)
	return p, nil
}

func (s *ProductService) UpdateProduct(id uint, p *model.Product) (*model.Product, error) {
	sprod := model.Product{}
	s.DB.Find(&sprod, id)
	s.DB.Model(&sprod).Updates(p)
	return &sprod, nil
}

func (s *ProductService) DeleteEquipo(id uint) error {
	s.DB.Delete(&model.Product{}, id)
	return nil
}
