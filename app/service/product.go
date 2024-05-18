package service

import (
	"onlineshop/app/repository"
	"onlineshop/models"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product models.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]models.GetProducts, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Delete(productId int) error {
	return s.repo.Delete(productId)
}

func (s *ProductService) Update(productId int, input models.UpdateProduct) error {
	return s.repo.Update(productId, input)
}

func (s *ProductService) GetById(productId int) (models.GetProduct, []models.GetProducts, error) {
	return s.repo.GetById(productId)
}
