package service

import (
	"errors"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/repository"
	"warehouse/logger"
)

func GetAllBatches() ([]models.Batch, error) {
	batches, err := repository.GetAllBatches()
	if err != nil {
		return []models.Batch{}, err
	}

	return batches, nil
}

func GetBatchByID(id int) (models.Batch, error) {
	batch, err := repository.GetBatchByID(id)
	if err != nil {
		return models.Batch{}, err
	}

	return batch, nil
}

func CreateBatch(b models.Batch) error {
	s, err := repository.GetStorageByAdressCode(b.AdressCode)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return err
	}

	if err == nil &&
		(s.Article != b.Article ||
			(s.Article == b.Article && b.Type == "out" && b.Quantity > s.Quantity)) {
		logger.Error.
			Printf("[service] CreateBatch(): article don't match or there are't enough quantity when out: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("article don't match or there are't enough quantity when out"), errs.ErrBadRequestBody)
		return err
	}

	if errors.Is(err, errs.ErrNotFound) {
		if b.Type == "out" {
			logger.Error.
				Printf("[service] CreateBatch(): there are't enough quantity when out: %s\n", errs.ErrBadRequestBody.Error())
			err := errors.Join(errors.New("there are't enough quantity when out"), errs.ErrBadRequestBody)
			return err
		}
		s = models.Storage{
			AdressCode: b.AdressCode,
			Article:    b.Article,
			Quantity:   b.Quantity,
		}
		_, err = repository.CreateStorage(s)
		if err != nil {
			return err
		}
	} else {
		if b.Type == "in" {
			s.Quantity += b.Quantity
		} else {
			s.Quantity -= b.Quantity
		}
		err = repository.UpdateStorageByAdressCode(s.AdressCode, s)
		if err != nil {
			return err
		}

		if s.Quantity == 0 {
			err = repository.DeleteStorageByAdressCode(s.AdressCode)
			if err != nil {
				return err
			}
		}
	}

	err = repository.CreateBatch(b)
	if err != nil {
		return err
	}

	return nil
}
