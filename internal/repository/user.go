package repository

import (
	"warehouse/internal/db"
	"warehouse/internal/models"
	"warehouse/logger"
)

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id,
					   username,
					   role,
					   full_name,
					   active
				FROM users
				WHERE username = $1;`, username)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByUsername(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func GetUserByID(id int) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id,
					   username,
					   role,
					   full_name,
					   active
				FROM users 
				WHERE id = $1;`, id)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByID(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func GetPasswordByID(id int) (password string) {
	_ = db.GetDBConn().Get(&password, `SELECT password
				FROM users 
				WHERE id = $1;`, id)
	return
}

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Select(&users, `SELECT id,
					   username,
					   role,
					   full_name,
					   active
				FROM users;`)
	if err != nil {
		logger.Error.
			Printf("[repository] GetAllUsers(): error duriing getting from database: %s\n", err.Error())
		return []models.User{}, translateError(err)
	}

	return users, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Get(&user, `SELECT id, 
					   username, 
					   role, 
					   full_name,
					   active
				FROM users 
				WHERE username = $1 AND password = $2;`, username, password)
	if err != nil {
		logger.Error.
			Printf("[repository] GetUserByUsernameAndPassword(): error duriing getting from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return user, nil
}

func CreateUser(u models.User) (models.User, error) {
	_, err := db.GetDBConn().Exec(`
			INSERT INTO users (username, password, role, full_name)
			VALUES ($1, $2, $3, $4);`, u.Username, u.Password, u.Role, u.FullName)
	if err != nil {
		logger.Error.
			Printf("[repository] CreateUser(): error duriing creating user from database: %s\n", err.Error())
		return models.User{}, translateError(err)
	}

	return GetUserByUsername(u.Username)
}

func UpdateUserByID(id int, u models.User) error {
	_, err := db.GetDBConn().Exec(`
			UPDATE users
			SET username = $1, password = $2, role = $3, full_name = $4, active = $5
			WHERE id = $6;`, u.Username, u.Password, u.Role, u.FullName, u.Active, id)
	if err != nil {
		logger.Error.
			Printf("[repository] UpdateUserByID(): error while updating data: %s\n", err.Error())
		return translateError(err)
	}
	return nil
}

// {
//     "username":"admin",
//     "role":"admin",
//     "full_name":"administrator admin",
//     "active":true
// }
