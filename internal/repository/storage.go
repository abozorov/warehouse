package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetStorageByAdressCode(adressCode string) (storage models.Storage, err error) {
	if err = db.GetDBConn().Get(&storage, `SELECT c.zone,
					   c.row,
					   s.adress_code,
					   s.article,
					   p.name,
					   p.price,
					   s.quantity
				FROM storages s, cells c, products p
				WHERE s.adress_code = $1 AND c.adress_code = $2 AND s.article = p.article;`, adressCode, adressCode); err != nil {
		logger.Error.
			Printf("[repository] GetStorageByAdressCode(): error duriing getting from database: %s\n", err.Error())
		return models.Storage{}, translateError(err)
	}

	return storage, nil
}

func GetAllStorages() (storage []models.Storage, err error) {
	err = db.GetDBConn().Select(&storage, `SELECT c.zone,
					   c.row,
					   s.adress_code,
					   s.article,
					   p.name,
					   p.price,
					   s.quantity
				FROM storages s, cells c, products p
				WHERE s.adress_code = c.adress_code AND s.article = p.article;`)
	if err != nil {
		logger.Error.
			Printf("[repository] GetAllStorages(): error duriing getting from database: %s\n", err.Error())
		return []models.Storage{}, translateError(err)
	}

	return storage, nil
}

func CreateStorage(storage models.Storage) (models.Storage, error) {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO Storages (adress_code, article, quantity)
			VALUES ($1, $2, $3);`, storage.AdressCode, storage.Article, storage.Quantity)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateStorage(): error duriing creating Storage from database: %s\n", err.Error())
		return models.Storage{}, translateError(err)
	}

	return GetStorageByAdressCode(storage.AdressCode)
}

func UpdateStorageByAdressCode(adressCode string, storage models.Storage) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE storages
			SET article = $1, quantity = $2
			WHERE adress_code = $3;`, storage.Article, storage.Quantity, adressCode)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateStorageByAdressCode(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteStorageByAdressCode(adressCode string) error {
	_, err := db.GetDBConn().Exec(`
			DELETE FROM storages
			WHERE adress_code = $1;`, adressCode)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateStorageByAdressCode(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
