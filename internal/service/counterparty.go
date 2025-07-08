package service

import (
	"errors"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/repository"
)

func CreateCounterparty(cp models.Counterparty) (models.Counterparty, error) {
	// Проверить существует ли контрагент
	_, err := repository.GetCounterpartyByEmail(cp.Email)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return models.Counterparty{}, err
	} else if err == nil {
		return models.Counterparty{}, errs.ErrCounterpartyAlreadyExists
	}

	// Создаем
	counterparty, err := repository.CreateCounterparty(cp)
	if err != nil {
		return models.Counterparty{}, err
	}

	return counterparty, nil
}

func GetcounterpartyByID(id int) (cp models.Counterparty, err error) {
	// достаем данные с бд
	cp, err = repository.GetCounterpartyByID(id)
	if err != nil {
		return models.Counterparty{}, err
	}

	// ответ
	return cp, nil
}

func GetAllCounterparties() (cp []models.Counterparty, err error) {
	// достаем данные с бд
	cp, err = repository.GetAllCounterparties()
	if err != nil {
		return []models.Counterparty{}, err
	}

	// ответ
	return cp, nil
}

func newCounterparty(newCp, oldCp models.Counterparty) models.Counterparty {
	if newCp.Name == "" {
		newCp.Name = oldCp.Name
	}

	if newCp.Contact == "" {
		newCp.Contact = oldCp.Contact
	}

	if newCp.Phone == "" {
		newCp.Phone = oldCp.Phone
	}

	if newCp.Email == "" {
		newCp.Email = oldCp.Email
	}
	return newCp
}

func UpdateCounterpartyByID(id int, cp models.Counterparty) error {
	// проверяем есть ли контрагенг с таким id
	originCp, err := repository.GetCounterpartyByID(id)
	if err != nil {
		return err
	}

	// обновляем данне
	if err = repository.UpdateCounterpartyByID(id, newCounterparty(cp, originCp)); err != nil {
		return err
	}
	return nil
}
