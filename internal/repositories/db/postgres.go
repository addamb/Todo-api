package db

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Postgres interface {
	PostgresConn() *sqlx.DB
}

func NewPostgresConn(url string) *sqlx.DB {
	dbx, err := sqlx.Connect("pgx", url) //make env maybe?
	if err != nil {
		log.Panicf("could not connect to DB: %v", err)
	}

	return dbx
}
