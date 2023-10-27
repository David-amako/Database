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

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	db, err := sql.Open("mysql", "root:*password*@tcp(localhost:3306)/nea_db")

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

	e.GET("/useraccounts", getuserAccount)

	e.Logger.Fatal(e.Start(":1323"))
}

func getuserAccount(c echo.Context) error {
	db, err := sql.Open("mysql", "root:*password*@tcp(localhost:3306)/nea_db")

	requested_id := c.Param("id")
	fmt.Println(requested_id)
	var firstname string
	var surname string
	var id string
	var email string
	var password string
	var registation_date string
	var address string
	var phone string

	err = db.QueryRow("SELECT * FROM nea_db.useraccounts WHERE user_id = 3;").Scan(&id, &firstname, &surname, &email,
		&password, &registation_date, &address, &phone)

	if err != nil {
		fmt.Println(err)
	}

	response := Useraccount{Id: id, Firstname: firstname, Surname: surname, Email: email, Password: password,
		Registation_date: registation_date, Address: address, Phone: phone}

	return c.JSON(http.StatusOK, response)
}

func createUserAccount(c echo.Context) error {
	emp := new(Useraccount)
	if err := c.Bind(emp); err != nil {
		return err
	}

	db := c.Get("db").(*sql.DB)

	// Insert data into the database
	// Example MySQL insert statement:
	insertStmt := "INSERT INTO `nea_db`.`useraccounts` (`Firstname`, `Surname`, `Email`, `Password`, `Registration_date`, `Address`, `Phone`) VALUES ('?', '?', '?', '?', '2023-10-27', '?', '?');"
	stmt, err := db.Prepare(insertStmt)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(emp.Firstname, emp.Surname, emp.Email, emp.Password, emp.Registation_date, emp.Address, emp.Phone)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result.LastInsertId())

	return c.JSON(http.StatusCreated, emp.Firstname)
}

func updateUserAccount(c echo.Context) error {
	id := c.Param("id") // Get the user ID from the URL parameter
	userAccount := new(Useraccount)
	if err := c.Bind(userAccount); err != nil {
		return err
	}

	db := c.Get("db").(*sql.DB)

	// Update data in the database
	// Example MySQL update statement:
	updateStmt := "UPDATE useraccounts SET Firstname = ?, Surname = ?, Email = ?, Password = ?, Registration_date = ?, Address = ?, Phone = ? WHERE User_id = ?"
	_, err := db.Exec(updateStmt, userAccount.Firstname, userAccount.Surname, userAccount.Email, userAccount.Password, userAccount.Registation_date, userAccount.Address, userAccount.Phone, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userAccount)
}

func deleteUserAccount(c echo.Context) error {
	id := c.Param("id") // Get the user ID from the URL parameter

	db := c.Get("db").(*sql.DB)

	// Delete the user account from the database
	// Example MySQL delete statement:
	deleteStmt := "DELETE FROM useraccounts WHERE User_id = ?"
	_, err := db.Exec(deleteStmt, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
