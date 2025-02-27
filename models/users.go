package models

import (
	"context"
	"playground/storage"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Create(user *User) error {
	db := storage.GetDB()
	args := pgx.NamedArgs{
		"userEmail":    user.Email,
		"userPassword": user.Password,
	}
	_, err := db.Exec(context.Background(), "INSERT INTO users (email, password) VALUES (@userEmail, @userPassword)", args)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindAll() ([]User, error) {
	db := storage.GetDB()
	rows, err := db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Password)
		res = append(res, user)
	}
	return res, nil
}
