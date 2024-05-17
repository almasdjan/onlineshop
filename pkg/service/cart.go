package service

import (
	"onlineshop/models"
	"onlineshop/pkg/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) Add(userId int, productId int) (int, error) {
	return s.repo.Add(userId, productId)
}

func (s *CartService) Minus(userId, productId int) error {
	return s.repo.Minus(userId, productId)
}

func (s *CartService) Plus(userId, productId int) error {
	return s.repo.Plus(userId, productId)
}

func (s *CartService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *CartService) GetAllFromCart(userId int) ([]models.GetProductsFromCart, error) {
	return s.repo.GetAllFromCart(userId)
}

func (s *CartService) GetTotalAmout(userId int) (float64, error) {
	return s.repo.GetTotalAmout(userId)
}
