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
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type addUseraccount struct {
	AddFirstname string `json:"Firstname"`
	AddSurname   string `json:"Surname"`
	AddEmail     string `json:"Email"`
	AddPassword  string `json:"Password"`
	AddAddress   string `json:"Address"`
	AddPhone     string `json:"Phone"`
}

type Useraccounts struct {
	Useraccounts map[string]Useraccount `json:"useraccounts"`
}
type Item struct {
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	Current_bid  string `json:"Current_bid"`
	Bid_duration string `json:"bid_duration"`
	Seller       string `json:"Seller"`
}
type Items struct {
	Items map[string]Item `json:"items"`
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

		var id string
		var email string
		var password string

		err = (query).Scan(&id, &email, &password)

		if err != nil {
			fmt.Println(err)
		}

		response.Useraccounts[id] = Useraccount{Email: email, Password: password}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforshoes(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Shoes\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforclothing(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Clothing\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforelectronics(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Electronics\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforsports(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Sports\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforhealth(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Health\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforfurniture(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Furniture\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforother(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, Bid_duration, Seller, Category FROM items WHERE category = \"Other\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var bid_duration string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &bid_duration, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, Bid_duration: bid_duration, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
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
	e.GET("/shoeitems", getiteminfoforshoes)
	e.GET("/clothingitems", getiteminfoforclothing)
	e.GET("/electronicsitems", getiteminfoforelectronics)
	e.GET("/sportsitems", getiteminfoforsports)
	e.GET("/healthitems", getiteminfoforhealth)
	e.GET("/furnitureitems", getiteminfoforfurniture)
	e.GET("/otheritems", getiteminfoforother)

	e.Logger.Fatal(e.Start(":1323"))
}
