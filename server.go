package main

import (
	"database/sql"
	"net/http"
	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Accounts struct {
	Id               int    `json:"User_id"`
	Firstname        string `json:"Firstname"`
	Surname          string `json:"Surname"`
	Email            string `json:"Email"`
	Password         string `json:"Password"`
	Registation_date string `json:"Registration_date"`
	Address          string `json:"Address"`
	Phone            string `json:"Phone"`
}

type UserData struct {
	Users map[string]Accounts `json:"Users"`
}

func getUser(c echo.Context, db *sql.DB) error {
	query, err := db.Query(`SELECT * FROM useraccounts`)

	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	data := UserData{Users: map[string]Accounts{}}

	for query.Next() {
		var name string

		var User_id int

		var firstname, surname, email, password, registation_date, address, phone string

		err := query.Scan(&name)

		if err != nil {
			panic(err.Error)
		}

		data.Users[name] = Accounts{Id: User_id, Firstname: firstname, Surname: surname, Email: email, Password: password, Registation_date: registation_date, Address: address, Phone: phone}
	}

	return c.JSON(http.StatusOK, data)
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
