package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetAllBatches() (batches []models.Batch, err error) {

	err = db.GetDBConn().Select(&batches, `
					SELECT b.id, b.date, b.type, b.counterparty_name, b.article, b.quantity, b.adress_code,
						cp.contact, cp.phone, cp.email,
						p.name AS product_name, p.price,
						c.zone, c.row,
						u.username, u.role, u.full_name, u.active
					FROM batches b
					INNER JOIN counterparties cp ON b.counterparty_name = cp.name
					INNER JOIN products p ON b.article = p.article
					INNER JOIN storages s ON b.adress_code = s.adress_code
					INNER JOIN cells c ON  b.adress_code = c.adress_code
					INNER JOIN users u ON b.username = u.username;`)

	if err != nil {
		logger.Error.
			Printf("[repository] GetAllBatches(): error duriing getting from database: %s\n", err.Error())
		return []models.Batch{}, translateError(err)
	}

	return batches, nil
}

func GetBatchByID(id int) (batch models.Batch, err error) {

	err = db.GetDBConn().Get(&batch, `
					SELECT b.id, b.date, b.type, b.counterparty_name, b.article, b.quantity, b.adress_code,
						cp.contact, cp.phone, cp.email,
						p.name AS product_name, p.price,
						c.zone, c.row,
						u.username, u.role, u.full_name, u.active
					FROM batches b
					INNER JOIN counterparties cp ON b.counterparty_name = cp.name
					INNER JOIN products p ON b.article = p.article
					INNER JOIN storages s ON b.adress_code = s.adress_code
					INNER JOIN cells c ON  b.adress_code = c.adress_code
					INNER JOIN users u ON b.username = u.username
					WHERE b.id = $1;`, id)

	if err != nil {
		logger.Error.
			Printf("[repository] GetBatchByID(): error duriing getting from database: %s\n", err.Error())
		return models.Batch{}, translateError(err)
	}

	return batch, nil
}

func GetFilteredBatches(table string, filter models.Filter) (batches []models.Batch, err error) {
	sqlQuery := `SELECT b.id, b.date, b.type, b.counterparty_name, b.article, b.quantity, b.adress_code,
					cp.contact, cp.phone, cp.email,
					p.name AS product_name, p.price,
					c.zone, c.row,
					u.username, u.role, u.full_name, u.active
				FROM batches b
					INNER JOIN counterparties cp ON b.counterparty_name = cp.name
					INNER JOIN products p ON b.article = p.article
					INNER JOIN storages s ON b.adress_code = s.adress_code
					INNER JOIN cells c ON  b.adress_code = c.adress_code
					INNER JOIN users u ON b.username = u.username
				WHERE $1 <= b.date AND b.date <= $2`

	// если есть id
	if filter.ID != "" {

		if filter.BatchType != "all" {
			table = " AND " + table + ".id = $3 "
			err = db.GetDBConn().Select(&batches, sqlQuery+table+" AND b.type = $4;",
				filter.DateFrom, filter.DateTo, filter.ID, filter.BatchType)
		} else {
			table = " AND " + table + ".id = $3 "
			err = db.GetDBConn().Select(&batches, sqlQuery+table+";",
				filter.DateFrom, filter.DateTo, filter.ID)
		}

	} else { // если нет id

		if filter.BatchType != "all" {
			err = db.GetDBConn().Select(&batches, sqlQuery+" AND b.type = $3;",
				filter.DateFrom, filter.DateTo, filter.BatchType)
		} else {
			err = db.GetDBConn().Select(&batches, sqlQuery+";",
				filter.DateFrom, filter.DateTo)
		}
	}

	if err != nil {
		logger.Error.
			Printf("[repository] GetFiltededBatches(): error duriing getting from database: %s\n", err.Error())
		return []models.Batch{}, translateError(err)
	}

	return batches, nil
}

func CreateBatch(b models.Batch) error {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO batches(type, counterparty_name, article, quantity, adress_code, username)
			values($1, $2, $3, $4, $5, $6);`, b.Type, b.CounterpartyName, b.Article, b.Quantity, b.AdressCode, b.Username)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateBatch(): error duriing creating batch from database: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}
