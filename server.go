package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Useraccount struct {
	Id               string `json:"User_id"`
	Firstname        string `json:"Firstname"`
	Surname          string `json:"Surname"`
	Email            string `json:"Email"`
	Password         string `json:"Password"`
	Registation_date string `json:"Registration_date"`
	Address          string `json:"Address"`
	Phone            string `json:"Phone"`
}

type Useraccounts struct {
	Useraccounts []Useraccount `json:"useraccounts"`
}

func getUser(c echo.Context, db *sql.DB) error {
	query, err := db.Query(`SELECT * FROM useraccounts`)

	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	for query.Next() {

	}

	return c.JSON(http.StatusOK, db)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	if err != nil {
		panic(err.Error)
	} else {
		fmt.Println("db is connected")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	e.GET("/useraccounts", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id)
		fmt.Println("^^")
		var firstname string
		var surname string
		var id string
		var email string
		var password string
		var registation_date string
		var addess string
		var phone string

		err = db.QueryRow("SELECT * FROM nea_db.useraccounts;").Scan(&id, &firstname, &surname, &email, &password, &registation_date, &addess, &phone)

		if err != nil {
			fmt.Println(err)
		}

		response := Useraccount{Id: id, Firstname: firstname, Surname: surname, Email: email, Password: password, Registation_date: registation_date, Address: addess, Phone: phone}
		return c.JSON(http.StatusOK, response)
	})

	/*e.GET("/", func(c echo.Context) error {
		return getUser(c, db)
	})*/

	e.Logger.Fatal(e.Start(":1323"))
}
