package repository

import (
	"fmt"
	"onlineshop/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (r *CartPostgres) Add(userId int, productId int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, product_id) values($1, $2) RETURNING id", cartsTable)
	row := r.db.QueryRow(query, userId, productId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CartPostgres) Minus(userId, productId int) error {

	row := fmt.Sprintf("SELECT quantity from %s  WHERE user_id =$1 AND product_id =$2",
		cartsTable)
	var quantity int
	err := r.db.Get(&quantity, row, userId, productId)
	logrus.Print(quantity)
	if err != nil {
		return err
	}

	if quantity == 1 {
		query := fmt.Sprintf("DELETE FROM %s WHERE user_id =$1 AND product_id =$2",
			cartsTable)

		_, err = r.db.Exec(query, userId, productId)
		if err != nil {
			return err
		}
		return nil
	}

	query := fmt.Sprintf("UPDATE %s SET quantity = quantity - 1 WHERE user_id =$1 AND product_id =$2",
		cartsTable)

	_, err = r.db.Exec(query, userId, productId)
	if err != nil {
		return err
	}
	return nil

}

func (r *CartPostgres) Plus(userId, productId int) error {
	query := fmt.Sprintf("UPDATE %s SET quantity = quantity + 1 WHERE user_id =$1 AND product_id =$2",
		cartsTable)

	_, err := r.db.Exec(query, userId, productId)
	if err != nil {
		return err
	}
	return nil

}

func (r *CartPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", cartsTable)
	_, err := r.db.Exec(query, userId)
	return err

}

func (r *CartPostgres) GetAllFromCart(userId int) ([]models.GetProductsFromCart, error) {
	var products []models.GetProductsFromCart
	query := fmt.Sprintf("SELECT p.id,p.name,p.image,p.price, c.quantity FROM %s p JOIN %s c ON p.id =c.product_id WHERE user_id = $1", productsTable, cartsTable)
	err := r.db.Select(&products, query, userId)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (r *CartPostgres) GetTotalAmout(userId int) (float64, error) {
	var totalAmount float64
	query := fmt.Sprintf("SELECT SUM(p.price *c.quantity) FROM %s p JOIN %s c ON p.id =c.product_id WHERE user_id = $1", productsTable, cartsTable)
	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&totalAmount); err != nil {
		return 0, err
	}
	return totalAmount, nil
}
