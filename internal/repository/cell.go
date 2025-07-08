package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetCellByAdressCode(adressCode string) (cell models.Cell, err error) {
	if err = db.GetDBConn().Get(&cell, `SELECT id,
					   zone,
					   row,
					   adress_code
				FROM cells
				WHERE adress_code = $1;`, adressCode); err != nil {
		logger.Error.
			Printf("[repository] GetCellByAdressCode(): error duriing getting from database: %s\n", err.Error())
		return models.Cell{}, translateError(err)
	}

	return cell, nil
}
 
func CreateCell(cell models.Cell) (models.Cell, error) {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO cells (zone, row, adress_code)
			VALUES ($1, $2, $3);`, cell.Zone, cell.Row, cell.AdressCode)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateCell(), cells: error duriing creating Cell from database: %s\n", err.Error())
		return models.Cell{}, translateError(err)
	}
	
	return GetCellByAdressCode(cell.AdressCode)
}

func GetCellByID(id int) (cell models.Cell, err error) {
	if err = db.GetDBConn().Get(&cell, `SELECT id,
					   zone,
					   row,
					   adress_code
				FROM cells
				WHERE id = $1;`, id); err != nil {
		logger.Error.
			Printf("[repository] GetCellByID(): error duriing getting from database: %s\n", err.Error())
		return models.Cell{}, translateError(err)
	}

	return cell, nil
}

func GetAllCells() (cells []models.Cell, err error) {
	err = db.GetDBConn().Select(&cells, `SELECT id,
					   zone,
					   row,
					   adress_code
				FROM cells;`)
	if err != nil {
		logger.Error.
			Printf("[repository] GetAllCells(): error duriing getting from database: %s\n", err.Error())
		return []models.Cell{}, translateError(err)
	}

	return cells, nil
}

func UpdateCellByID(id int, cell models.Cell) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE cells
			SET zone = $1, row = $2, adress_code = $3
			WHERE id = $4;`, cell.Zone, cell.Row, cell.AdressCode, id)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateCellByID(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
