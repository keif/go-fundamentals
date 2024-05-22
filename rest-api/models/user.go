package models

import "example.com/rest-api/db"

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
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

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = uint(userId)
	return err
}
