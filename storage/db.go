package storage

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func InitDB() {
	// link example postgres://user:pass@localhost:5432/dbName
	conn, err := pgx.Connect(context.Background(), os.Args[1])
	if err != nil {
		fmt.Println("error with connecting to db", os.Args[1])
		os.Exit(1)
	}
	Db = conn
}

func GetDB() *pgx.Conn {
	return Db
}

func Migrate() error {

	text, err := getFileText()
	if err != nil {
		return err
	}

	sqlQueries := getSqlQueries(&text)
	for _, query := range sqlQueries {
		rows, err := Db.Query(context.Background(), query)
		if err != nil {
			return err
		}
		rows.Close()
	}

	return nil
}

func getSqlQueries(text *string) []string {
	sqlQueries := []string{}
	temp := []string{}
	for _, letter := range *text {
		if string(letter) != ";" {
			temp = append(temp, string(letter))
		} else {
			temp = append(temp, ";")
			sqlQueries = append(sqlQueries, strings.Join(temp, ""))
			temp = []string{}
		}
	}

	return sqlQueries
}

func getFileText() (string, error) {
	bytes, err := os.ReadFile("db.sql")
	if err != nil {
		return "", err
	}
	text := string(bytes)
	return text, nil
}
