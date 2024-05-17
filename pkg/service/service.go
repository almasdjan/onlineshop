package service

import (
	"onlineshop/models"
	"onlineshop/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	IsAdmin(userId int) (bool, error)
}

type Product interface {
	Create(product models.Product) (int, error)
	GetAll() ([]models.GetProducts, error)
	Delete(productId int) error
	Update(productId int, input models.UpdateProduct) error
	GetById(productId int) (models.GetProduct, []models.GetProducts, error)
}

type Cart interface {
	Add(userId int, productId int) (int, error)
	Minus(userId, productId int) error
	Plus(userId, productId int) error
	Delete(userId int) error
	GetAllFromCart(userId int) ([]models.GetProductsFromCart, error)
	GetTotalAmout(userId int) (float64, error)
}

type Order interface {
	Create(userId int, totalAmount float64, address models.Address, products []models.GetProductsFromCart) (int, error)
	GetAllForUser(userId int) ([]models.OrderForUser, error)
	GetAll() ([]models.OrderForAdmin, error)
	Update(orderId int, statusId models.Status) error
}

type Service struct {
	Authorization
	Product
	Cart
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
		Cart:          NewCartService(repos.Cart),
		Order:         NewOrderService(repos.Order),
	}
}
