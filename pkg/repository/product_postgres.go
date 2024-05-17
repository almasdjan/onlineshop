package repository

import (
	"fmt"
	"onlineshop/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product models.Product) (int, error) {
	var recommendedProductss []string
	for _, v := range product.RecommendedProducts {
		recommendedProductss = strings.Split(v, ",")

	}

	var recommendedProductsArray []int
	for _, v := range recommendedProductss {
		pId, err := strconv.Atoi(v)
		if err != nil {

			return 0, err
		}
		recommendedProductsArray = append(recommendedProductsArray, pId)
	}

	var id int
	productQuery := fmt.Sprintf("INSERT INTO %s (name, image, price,height, size, instruction, description) values($1, $2,$3,$4,$5,$6,$7) RETURNING id", productsTable)
	row := r.db.QueryRow(productQuery, product.Name, product.Image, product.Price, product.Height, product.Size, product.Instruction, product.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	logrus.Print(recommendedProductsArray)
	for _, rProduct := range recommendedProductsArray {

		recommendedQuery := fmt.Sprintf("INSERT INTO %s (product_id, recommended_product) values ($1,$2)", recommendedProductsTable)
		_, err := r.db.Exec(recommendedQuery, id, rProduct)
		if err != nil {
			return id, err
		}
	}
	return id, nil

}

func (r *ProductPostgres) GetAll() ([]models.GetProducts, error) {
	var products []models.GetProducts
	query := fmt.Sprintf("SELECT id, name, image, price FROM %s order by id", productsTable)
	logrus.Print(query)
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, err
}

func (r *ProductPostgres) Delete(productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		productsTable)
	_, err := r.db.Exec(query, productId)
	return err
}

func (r *ProductPostgres) Update(productId int, input models.UpdateProduct) error {
	var recommendedProductss []string
	for _, v := range input.RecommendedProducts {
		recommendedProductss = strings.Split(v, ",")

	}

	var recommendedProductsArray []int
	for _, v := range recommendedProductss {
		pId, err := strconv.Atoi(v)
		if err != nil {

			return err
		}
		recommendedProductsArray = append(recommendedProductsArray, pId)
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Image != "" {
		setValues = append(setValues, fmt.Sprintf("image = $%d", argId))
		args = append(args, input.Image)
		argId++
	}

	if input.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price = $%d", argId))
		args = append(args, input.Price)
		argId++
	}

	if input.Height != "" {
		setValues = append(setValues, fmt.Sprintf("height = $%d", argId))
		args = append(args, input.Height)
		argId++
	}

	if input.Size != "" {
		setValues = append(setValues, fmt.Sprintf("size = $%d", argId))
		args = append(args, input.Size)
		argId++
	}

	if input.Instruction != "" {
		setValues = append(setValues, fmt.Sprintf("instruction = $%d", argId))
		args = append(args, input.Instruction)
		argId++
	}

	if input.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id =$%d ",
		productsTable, setQuery, argId)

	args = append(args, productId)

	logrus.Print(setQuery)
	logrus.Print(query)
	logrus.Print(args)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if recommendedProductsArray != nil {
		query := fmt.Sprintf("DELETE FROM %s WHERE product_id = $1", recommendedProductsTable)
		_, err := r.db.Exec(query, productId)
		if err != nil {
			return err
		}
		for _, rProduct := range recommendedProductsArray {
			recommendedQuery := fmt.Sprintf("INSERT INTO %s (product_id, recommended_product) values ($1,$2)", recommendedProductsTable)
			_, err = r.db.Exec(recommendedQuery, productId, rProduct)
			if err != nil {
				return err
			}
		}
	}
	return err

}

func (r *ProductPostgres) GetById(productId int) (models.GetProduct, []models.GetProducts, error) {
	var product models.GetProduct
	query := fmt.Sprintf("SELECT id, name, image, price, height, size, instruction,description FROM %s WHERE id = $1", productsTable)
	logrus.Print(query)
	err := r.db.Get(&product, query, productId)
	if err != nil {
		return product, nil, err
	}

	var products []models.GetProducts
	rQuery := fmt.Sprintf("SELECT p.id, p.name, p.image, p.price FROM %s p JOIN %s rp ON p.id = rp.product_id WHERE  p.id=$1 AND rp.product_id=$1 ", productsTable, recommendedProductsTable)
	logrus.Print(rQuery)
	err = r.db.Select(&products, rQuery, productId)
	if err != nil {
		return product, nil, err
	}

	return product, products, err
}
