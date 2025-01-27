package models

import (
	"database/sql"
	"errors"
	"fmt"

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
		return fmt.Errorf("failed to prepare user save query: %w", err)
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return fmt.Errorf("failed to execute user save query: %w", err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last insert ID: %w", err)
	}
	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
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
