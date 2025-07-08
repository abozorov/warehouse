package service

import (
	"warehouse/internal/models"
	"warehouse/internal/repository"
)

func GetStats(table string, filter models.Filter) ([]models.Batch, error) {
	batches, err := repository.GetFilteredBatches(table, filter)

	return batches, err
}
