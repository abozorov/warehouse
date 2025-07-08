package db

import "warehouse/logger"

func InitMigrations() error {
	userTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(60) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(60) NOT NULL,
		full_name VARCHAR(100) NOT NULL,
		active BOOL DEFAULT TRUE
	);`

	_, err := db.Exec(userTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create users table: %v", err.Error())
		return err
	}

	counterpartiesTableQuery := `
	CREATE TABLE IF NOT EXISTS counterparties (
		id SERIAL PRIMARY KEY,
		name VARCHAR(60) UNIQUE NOT NULL,
		contact VARCHAR(60),
		phone VARCHAR(15) UNIQUE NOT NULL,
		email VARCHAR(60) UNIQUE NOT NULL
	);`

	_, err = db.Exec(counterpartiesTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create counterparties table: %v", err.Error())
		return err
	}

	productsTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		article VARCHAR(30) UNIQUE NOT NULL,
		name VARCHAR(100),
		price REAL DEFAULT 0
	);`

	_, err = db.Exec(productsTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create products table: %v", err.Error())
		return err
	}

	cellsTableQuery := `
	CREATE TABLE IF NOT EXISTS cells (
		id SERIAL PRIMARY KEY,
		zone VARCHAR(60),
		row INTEGER,
		adress_code VARCHAR(60) UNIQUE NOT NULL
	);`

	_, err = db.Exec(cellsTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create cells table: %v", err.Error())
		return err
	}

	storagesTableQuery := `
	CREATE TABLE IF NOT EXISTS storages (
		adress_code VARCHAR(60) UNIQUE REFERENCES cells(adress_code) ON UPDATE CASCADE,
		article VARCHAR(30) REFERENCES products(article) ON UPDATE CASCADE DEFAULT '',
		quantity INTEGER DEFAULT 0
	);`

	_, err = db.Exec(storagesTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create storages table: %v", err.Error())
		return err
	}

	batchesTableQuery := `
	CREATE TABLE IF NOT EXISTS batches (
		id SERIAL PRIMARY KEY,
		date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		type VARCHAR(30) NOT NULL,
		counterparty_name VARCHAR(60) REFERENCES counterparties(name) ON UPDATE CASCADE,
		article VARCHAR(30) REFERENCES products(article) ON UPDATE CASCADE,
		quantity INTEGER,
		adress_code VARCHAR(60) REFERENCES cells(adress_code) ON UPDATE CASCADE,
		username VARCHAR(60) REFERENCES users(username) ON UPDATE CASCADE
	);`

	_, err = db.Exec(batchesTableQuery)
	if err != nil {
		logger.Error.Printf("[db] InitMigrations(): error during create batches table: %v", err.Error())
		return err
	}

	return nil
}
