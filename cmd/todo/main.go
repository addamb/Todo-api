package main

import (
	"fmt"
	"os"

	"github.com/addamb/todo-api/internal/handler"
	"github.com/addamb/todo-api/internal/repositories/db"
	"github.com/addamb/todo-api/internal/repositories/db/queries"
	"github.com/labstack/echo/v4"
)

const (
	defaultPort = "8080"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port = fmt.Sprintf(":%s", defaultPort)
	}

	url := "postgresql://postgres:nomad@localhost:5432/todo?sslmode=disable"

	db := db.NewPostgresConn(url)

	//Check db connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	queries := queries.NewSQLXQueries(db)

	h := handler.NewHandler(queries)

	e := echo.New()
	api := e.Group("/api")

	h.RegisterRoutes(api)

	e.Logger.Fatal(e.Start(port))
}
