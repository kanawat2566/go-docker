package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kanawat2566/go-docker/rest/handler"
	"github.com/labstack/echo"
)

func main() {
	connStr := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewApplication(db)

	e := echo.New()
	e.GET("/", h.Greeting)

	serverPort := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(serverPort))
}
