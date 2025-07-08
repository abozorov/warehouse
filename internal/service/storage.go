package service

import (
	"warehouse/internal/models"
	"warehouse/internal/repository"
)

func GetStorageByAdressCode(adressCode string) (models.Storage, error) {
	storage, err := repository.GetStorageByAdressCode(adressCode)
	if err != nil {
		return models.Storage{}, err
	}

	return storage, nil
}

func GetAllStorages() ([]models.Storage, error) {
	storage, err := repository.GetAllStorages()
	if err != nil {
		return []models.Storage{}, err
	}

	return storage, nil
}
