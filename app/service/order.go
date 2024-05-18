package service

import (
	"onlineshop/app/repository"
	"onlineshop/models"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(userId int, totalAmount float64, address models.Address, products []models.GetProductsFromCart) (int, error) {
	return s.repo.Create(userId, totalAmount, address, products)
}

func (s *OrderService) GetAllForUser(userId int) ([]models.OrderForUser, error) {
	return s.repo.GetAllForUser(userId)
}

func (s *OrderService) GetAll() ([]models.OrderForAdmin, error) {
	return s.repo.GetAll()
}

func (s *OrderService) Update(orderId int, statusId models.Status) error {
	return s.repo.Update(orderId, statusId)

}
