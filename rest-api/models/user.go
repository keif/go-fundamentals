package models

import (
	"database/sql"
	"errors"
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
	"fmt"
)

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `gorm:"size:255;unique" json:"email" binding:"required"`
	Password string `gorm:"size:255" json:"password" binding:"required"`
}

func (u User) Save() error {
	var err error
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("Credentials invalid.")
		}
		return errors.New("Database error.")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("Credentials invalid.")
		}
		return errors.New("Database error.")
	}

	return nil
}
