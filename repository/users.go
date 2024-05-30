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


func Login(db *sql.DB, dataUser structs.Users) (int64, error) {
	checkUserExist := "SELECT id, email, password FROM users WHERE email = $1"

	var user structs.Users
	data := db.QueryRow(checkUserExist, dataUser.Email).Scan(&user.ID, &user.Email, &user.Password)


	if data == sql.ErrNoRows {
		return 0, errors.New("Email tidak ditemukan")
	} 

	if user.Password != dataUser.Password {
		return 0, errors.New("Password Salah")
	}

	return user.ID, nil
}
