package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Postgres interface {
	PostgresConn() *sqlx.DB
}

func getInfo() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
}

func NewPostgresConn() *sqlx.DB {
	dbx, err := sqlx.Connect("pgx", getInfo())
	if err != nil {
		log.Panicf("could not connect to DB: %v", err)
	}

	return dbx
}
