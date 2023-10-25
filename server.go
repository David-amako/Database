package main

import (
	"database/sql"
	"net/http"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getUser(c echo.Context, db *sql.DB) error {
	query, err := db.Query(`SELECT * FROM useraccounts`)

	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	return c.JSON(http.StatusOK, db)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	db, err := sql.Open("mysql", "root:*password*@tcp(localhost:3306)/nea_db")

	if err != nil {
		panic(err.Error)
	}

	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return getUser(c, db)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
