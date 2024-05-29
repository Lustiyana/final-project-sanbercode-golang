package repository

import (
	"database/sql"
	"simple-social-media/structs"
	"errors"
	// "fmt"
)

func Register(db *sql.DB, dataUser structs.Users) error {
	checkUserExist := "SELECT email FROM users WHERE email = $1"

	var storedEmail string
	err := db.QueryRow(checkUserExist, dataUser.Email).Scan(&storedEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			sql := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3)"
	
			_, err = db.Exec(sql, dataUser.Email, dataUser.Name, dataUser.Password)
			if err != nil {
				return err
			}

			return nil
		} 
		return err
	}

	return errors.New("Email sudah terdaftar")
}


func Login(db *sql.DB, dataUser structs.Users) error {
	checkUserExist := "SELECT email, password FROM users WHERE email = $1"

	var user structs.Users
	data := db.QueryRow(checkUserExist, dataUser.Email).Scan(&user.Email, &user.Password)

	if data == sql.ErrNoRows {
		return errors.New("Email tidak ditemukan")
	} 

	if user.Password != dataUser.Password {
		return errors.New("Password Salah")
	}

	return nil
}
