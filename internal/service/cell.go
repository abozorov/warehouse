package service

import (
	"errors"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/repository"
)

func CreateCell(cell models.Cell) (models.Cell, error) {
	// Проверить есть ли такой продукт
	_, err := repository.GetCellByAdressCode(cell.AdressCode)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return models.Cell{}, err
	} else if err == nil {
		return models.Cell{}, errs.ErrCellAlreadyExists
	}

	// Добавить продукт
	cell, err = repository.CreateCell(cell)
	if err != nil {
		return models.Cell{}, err
	}

	// Ответ
	return cell, nil
}

func GetCellByID(id int) (models.Cell, error) {
	Cell, err := repository.GetCellByID(id)
	if err != nil {
		return models.Cell{}, err
	}

	return Cell, nil
}

func GetAllCells() ([]models.Cell, error) {
	Cells, err := repository.GetAllCells()
	if err != nil {
		return []models.Cell{}, err
	}

	return Cells, nil
}

func newCell(oldCell, newCell models.Cell) models.Cell {
	if newCell.AdressCode == "" {
		newCell.AdressCode = oldCell.AdressCode
	}

	if newCell.Row == 0 {
		newCell.Row = oldCell.Row
	}

	if newCell.Zone == "" {
		newCell.Zone = oldCell.Zone
	}

	return newCell
}

func UpdateCellByID(id int, p models.Cell) error {
	// проверяем есть ли товар с таким id
	cell, err := repository.GetCellByID(id)
	if err != nil {
		return err
	}

	// если есть то обновляем данные
	if err = repository.UpdateCellByID(id, newCell(cell, p)); err != nil {
		return err
	}
	return nil
}
