package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func UserSignup(user User) error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	return err
}

func UserLogin(user User) (error, bool, int) {
	var password string
	var id int

	query := "SELECT id, password FROM users WHERE email = ?"

	if err := db.DB.QueryRow(query, user.Email).Scan(&id, &password); err != nil {
		return errors.New("wrong credentials"), false, 0
	}

	comp := utils.CheckPasswordHash(user.Password, password)

	return nil, comp, id
}
