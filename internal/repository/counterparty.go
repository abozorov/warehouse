package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetCounterpartyByEmail(email string) (cp models.Counterparty, err error) {
	err = db.GetDBConn().Get(&cp, `SELECT id,
					name,
					contact,
					phone,
					email
				FROM counterparties
				WHERE email = $1;`, email)
	if err != nil {
		logger.Error.
			Printf("[repository] GetCounterpartyByEmail(): error duriing getting from database: %s\n", err.Error())
		return models.Counterparty{}, translateError(err)
	}

	return cp, nil
}

func CreateCounterparty(cp models.Counterparty) (models.Counterparty, error) {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO counterparties (name, contact, phone, email)
			VALUES ($1, $2, $3, $4);`, cp.Name, cp.Contact, cp.Phone, cp.Email)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateCounterparty(): error duriing creating counterparty from database: %s\n", err.Error())
		return models.Counterparty{}, translateError(err)
	}

	return GetCounterpartyByEmail(cp.Email)
}

func GetCounterpartyByID(id int) (cp models.Counterparty, err error) {
	err = db.GetDBConn().Get(&cp, `SELECT id,
					name,
					contact,
					phone,
					email
				FROM counterparties
				WHERE id = $1;`, id)
	if err != nil {
		logger.Error.
			Printf("[repository] GetCounterpartyByID(): error duriing getting from database: %s\n", err.Error())
		return models.Counterparty{}, translateError(err)
	}
	return cp, nil
}

func GetAllCounterparties() (cp []models.Counterparty, err error) {
	err = db.GetDBConn().Select(&cp, `SELECT id,
					name,
					contact,
					phone,
					email
				FROM counterparties;`)
	if err != nil {
		logger.Error.
			Printf("[repository] GetAllCounterparties(): error duriing getting from database: %s\n", err.Error())
		return []models.Counterparty{}, translateError(err)
	}
	return cp, nil
}

func UpdateCounterpartyByID(id int, cp models.Counterparty) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE counterparties
			SET name = $1, contact = $2, phone = $3, email = $4
			WHERE id = $5;`, cp.Name, cp.Contact, cp.Phone, cp.Email, id)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateCounterpartyByID(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
