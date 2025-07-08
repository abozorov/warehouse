package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetProductByArticle(article string) (p models.Product, err error) {
	if err = db.GetDBConn().Get(&p, `SELECT id,
					   article,
					   name,
					   price
				FROM products
				WHERE article = $1;`, article); err != nil {
		logger.Error.
			Printf("[repository] GetProductByArticle(): error duriing getting from database: %s\n", err.Error())
		return models.Product{}, translateError(err)
	}

	return p, nil
}

func CreateProduct(p models.Product) (models.Product, error) {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO products (article, name, price)
			VALUES ($1, $2, $3);`, p.Article, p.Name, p.Price)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateProduct(): error duriing creating product from database: %s\n", err.Error())
		return models.Product{}, translateError(err)
	}
	return GetProductByArticle(p.Article)
}

func GetProductByID(id int) (p models.Product, err error) {
	if err = db.GetDBConn().Get(&p, `SELECT id,
					   article,
					   name,
					   price
				FROM products
				WHERE id = $1;`, id); err != nil {
		logger.Error.
			Printf("[repository] GetProductByID(): error duriing getting from database: %s\n", err.Error())
		return models.Product{}, translateError(err)
	}

	return p, nil
}

func GetAllProducts() (products []models.Product, err error) {
	err = db.GetDBConn().Select(&products, `SELECT id,
					   article,
					   name,
					   price
				FROM products;`)
	if err != nil {
		logger.Error.
			Printf("[repository] GetAllProducts(): error duriing getting from database: %s\n", err.Error())
		return []models.Product{}, translateError(err)
	}

	return products, nil
}

func UpdateProductByID(id int, p models.Product) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE products
			SET article = $1, name = $2, price = $3
			WHERE id = $4;`, p.Article, p.Name, p.Price, id)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateProductByID(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
