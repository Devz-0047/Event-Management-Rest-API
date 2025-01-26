package models

import (
	"database/sql"
	"errors"

	"example.com/REST/db"
	"example.com/REST/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if err != nil {
		return err
	}
	hashPassword, err := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}
func (u *User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid email or password")
		}
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid email or password")
	}
	return nil
}
