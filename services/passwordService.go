package services

import (
	"database/sql"
	"fmt"
	databaseConn "password-generator/database"
)

func SearchAllPasswordsDatabase() ([]struct {
    Name     string
    Password string
}, error) {
	db, err := databaseConn.ConnectToDatabase()
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %v", err)
	}
	defer db.Close()

	query := "SELECT password_name, password FROM passwords"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve passwords: %v", err)
	}
	defer rows.Close()

	var passwords []struct {
		Name     string
		Password string
	}

	for rows.Next() {
		var passwordName, password string
		if err := rows.Scan(&passwordName, &password); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		passwords = append(passwords, struct {
			Name     string
			Password string
		}{
			Name:     passwordName,
			Password: password,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return passwords, nil
}


func SearchPasswordByNameDatabase(passwordName string) string {
	db, err := databaseConn.ConnectToDatabase()
	if err != nil {
		return fmt.Sprintf("Could not connect to database: %v", err)
	}
	defer db.Close()

	query := "SELECT password FROM passwords WHERE password_name = $1"
	row := db.QueryRow(query, passwordName)

	var password string
	err = row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Sprintf("No password found with name: %v", passwordName)
		}
		return fmt.Sprintf("Could not retrieve password: %v", err)
	}

	return password
}


func PasswordToDatabase(password string, passwordName string) error {
	db, err := databaseConn.ConnectToDatabase()
	if err != nil {
		return fmt.Errorf("Could not connect to database : %v", err)
	}
	defer db.Close()

	query := "INSERT INTO passwords (password_name, password) VALUES ($1, $2)"
	_, err = db.Exec(query, passwordName, password)
	if err != nil{
		return fmt.Errorf("could not insert password: %v", err)
	}

	return nil
}