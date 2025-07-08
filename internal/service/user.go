package service

import (
	"errors"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/repository"
	"warehouse/utils"
)

func CreateUser(u models.User) (models.User, error) {

	// 1. Проверить существует ли пользователь с таким username
	_, err := repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return models.User{}, err
	} else if err == nil {
		return models.User{}, errs.ErrUserAlreadyExists
	}

	// 2. Захешировать пароль
	u.Password = utils.GenerateHash(u.Password)

	// 3. Создаем пользователя
	user, err := repository.CreateUser(u)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	// 1. Хешируем пароль
	password = utils.GenerateHash(password)

	// 2. Отправляем запрос в бд
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrIncorrectUsernameOrPassword
		}
		return models.User{}, err
	}
	return user, nil
}

func GetUserByID(id int) (models.User, error) {
	// достаем данные с бд
	user, err := repository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	// ответ
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func newUser(oldU, newU models.User) models.User {
	if newU.Username == "" {
		newU.Username = oldU.Username
	}

	if newU.Password == "" {
		newU.Password = oldU.Password
	} else {
		newU.Password = utils.GenerateHash(newU.Password)
	}

	if newU.Role == "" {
		newU.Role = oldU.Role
	}

	if newU.FullName == "" {
		newU.FullName = oldU.FullName
	}

	return newU
}

func UpdateUserByID(id int, u models.User) error {

	// проверяем есть ли пользователь с таким id
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Password = repository.GetPasswordByID(id)

	// если есть то обновляем данные
	if err = repository.UpdateUserByID(id, newUser(user, u)); err != nil {
		return err
	}
	return nil
}
