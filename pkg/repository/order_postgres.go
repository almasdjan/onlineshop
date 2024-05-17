package repository

import (
	"fmt"
	"onlineshop/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(userId int, totalAmount float64, address models.Address, products []models.GetProductsFromCart) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, total_amount,address) values($1, $2,$3) RETURNING id", ordersTable)
	row := r.db.QueryRow(query, userId, totalAmount, address.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	for _, v := range products {
		var itemId int
		itemsQuery := fmt.Sprintf("INSERT INTO %s (order_id, product_id,quantity, price) values($1, $2,$3,$4) RETURNING id", ordersItems)
		row := r.db.QueryRow(itemsQuery, id, v.Id, v.Quantity, v.Price)
		if err := row.Scan(&itemId); err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *OrderPostgres) GetAllForUser(userId int) ([]models.OrderForUser, error) {
	var orders []models.OrderForUser
	query := fmt.Sprintf(`select o.id, o.total_amount, os.status, o.created_date, o.address, oi.product_id,p.name,p.image,p.price,oi.quantity 
						  FROM %s o JOIN %s oi ON o.id = oi.order_id  
						  JOIN %s os ON o.status_id = os.id
						  JOIN %s p ON oi.product_id = p.id
						  WHERE o.user_id = $1
						  ORDER BY o.id DESC`,
		ordersTable, ordersItems, orderStatusTable, productsTable)
	logrus.Print(query)
	err := r.db.Select(&orders, query, userId)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (r *OrderPostgres) GetAll() ([]models.OrderForAdmin, error) {
	var orders []models.OrderForAdmin
	query := fmt.Sprintf(`select o.id, o.total_amount, os.status, o.created_date,o.user_id, o.address, oi.product_id,p.name,p.image,p.price,oi.quantity 
						  FROM %s o JOIN %s oi ON o.id = oi.order_id  
						  JOIN %s os ON o.status_id = os.id
						  JOIN %s p ON oi.product_id = p.id
						  ORDER BY o.id DESC`,
		ordersTable, ordersItems, orderStatusTable, productsTable)
	logrus.Print(query)
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (r *OrderPostgres) Update(orderId int, statusId models.Status) error {
	query := fmt.Sprintf("UPDATE %s SET status_id = $2 where id = $1", ordersTable)
	_, err := r.db.Exec(query, orderId, statusId.StatusId)
	return err
}
