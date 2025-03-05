package storage

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Db *pgx.Conn

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_STR"))
	if err != nil {
		log.Fatal("error with connecting to db", err)
		os.Exit(1)
	}
	Db = conn
}

func GetDB() *pgx.Conn {
	return Db
}
