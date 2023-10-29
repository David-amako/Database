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
	//Id               string `json:"User_id"`
	//Firstname        string `json:"Firstname"`
	//Surname          string `json:"Surname"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	//Registation_date string `json:"Registration_date"`
	//Address          string `json:"Address"`
	//Phone            string `json:"Phone"`
}

type Useraccounts struct {
	Useraccounts map[string]Useraccount `json:"useraccounts"`
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

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db) // Set 'db' in the context
			return next(c)
		}
	})

	e.GET("/useraccounts", getuserinfoAccount)

	e.Logger.Fatal(e.Start(":1323"))
}

func getuserinfoAccount(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT user_id, email , Password FROM nea_db.useraccounts")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Useraccounts{Useraccounts: map[string]Useraccount{}}

	for query.Next() {

		//requested_id := c.Param("id")
		//fmt.Println(requested_id)
		//var firstname string
		//var surname string
		var id string
		var email string
		var password string
		//var registation_date string
		//var address string
		//var phone string

		err = (query).Scan(&id, &email, &password)

		if err != nil {
			fmt.Println(err)
		}

		response.Useraccounts[id] = Useraccount{Email: email, Password: password}

	}
	return c.JSON(http.StatusOK, response)
}
