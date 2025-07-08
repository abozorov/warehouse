package service

import (
	"errors"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/repository"
)

func CreateProduct(p models.Product) (models.Product, error) {
	// Проверить есть ли такой продукт
	_, err := repository.GetProductByArticle(p.Article)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return models.Product{}, err
	} else if err == nil {
		return models.Product{}, errs.ErrProductAlreadyExists
	}

	// Добавить продукт
	p, err = repository.CreateProduct(p)
	if err != nil {
		return models.Product{}, err
	}

	// Ответ
	return p, nil
}

func GetProductByID(id int) (models.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func newProduct(oldP, newP models.Product) models.Product {
	if newP.Article == "" {
		newP.Article = oldP.Article
	}

	return newP
}

func UpdateProductByID(id int, p models.Product) error {
	// проверяем есть ли товар с таким id
	product, err := repository.GetProductByID(id)
	if err != nil {
		return err
	}

	// если есть то обновляем данные
	if err = repository.UpdateProductByID(id, newProduct(product, p)); err != nil {
		return err
	}
	return nil
}
