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
	AddFirstname         string `json:"Firstname"`
	AddSurname           string `json:"Surname"`
	AddEmail             string `json:"Email"`
	AddPassword          string `json:"Password"`
	AddRegistration_date string `json:"Registration_date"`
	AddPhone             string `json:"Phone"`
}
type additem struct {
	AddTitle        string `json:"Title"`
	AddDescription  string `json:"Description"`
	AddStarting_bid string `json:"Starting_bid"`
	AddCurrent_bid  string `json:"Current_bid"`
	AddBid_duration string `json:"Bid_duration"`
	AddSeller       string `json:"Seller "`
	AddCategory     string `json:"Category"`
}
type UseraccountDetails struct {
	Firstname        string `json:"Firstname"`
	Surname          string `json:"Surname"`
	Email            string `json:"Email"`
	Password         string `json:"Password"`
	RegistrationDate string `json:"Registration_Date"`
	Address          string `json:"Address"`
	Phone            string `json:"Phone"`
}
type UseraccountsDetails struct {
	Useraccounts map[string]UseraccountDetails `json:"useraccounts"`
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

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Field       string `json:"field"`
	Value       string `json:"value"`
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
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func adduseracc(c echo.Context) error {
	// Parse request body
	var requestData addUseraccount
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Connect to the database
	db, err := connectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	// Perform database operations (INSERT, UPDATE, etc.) using requestData
	result, err := db.Exec("INSERT INTO `nea_db`.`useraccounts` (`Firstname`, `Surname`, `Email`, `Password`,`Registration_date`, `Phone`) VALUES (?, ?, ?, ?, ?, ?)",
		requestData.AddFirstname, requestData.AddSurname, requestData.AddEmail, requestData.AddPassword, requestData.AddRegistration_date, requestData.AddPhone)
	if err != nil {
		fmt.Println(err) // Handle the error appropriately
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert data"})
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err) // Handle the error appropriately
	}

	// Return a response
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Data inserted successfully", "insertedID": insertedID})
}
func additems(c echo.Context) error {
	// Parse request body
	var requestData additem
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Connect to the database
	db, err := connectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	// Perform database operations (INSERT, UPDATE, etc.) using requestData
	result, err := db.Exec("INSERT INTO `nea_db`.`items` (`Title`, `Description`, `Starting_bid`, `Current_bid`, `Bid_duration`, `Seller`, `Category`) VALUES (?, ?, ?, ?, ?, ?, ?);",
		requestData.AddTitle, requestData.AddDescription, requestData.AddStarting_bid, requestData.AddCurrent_bid, requestData.AddBid_duration, requestData.AddSeller, requestData.AddCategory)
	if err != nil {
		fmt.Println(err) // Handle the error appropriately
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert data"})
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err) // Handle the error appropriately
	}

	// Return a response
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Data inserted successfully", "insertedID": insertedID})
}
func getUserDetails(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	query, err := db.Query("SELECT user_id, Firstname, Surname, email, Password, Registration_date, Address, Phone FROM nea_db.useraccounts")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to execute query"})
	}
	defer query.Close()

	response := UseraccountsDetails{Useraccounts: map[string]UseraccountDetails{}}

	for query.Next() {
		var id string
		var firstname string
		var surname string
		var email string
		var password string
		var registrationDate string
		var address string
		var phone string

		err = query.Scan(&id, &firstname, &surname, &email, &password, &registrationDate, &address, &phone)
		if err != nil {
			fmt.Println(err)
		}

		response.Useraccounts[id] = UseraccountDetails{
			Firstname:        firstname,
			Surname:          surname,
			Email:            email,
			Password:         password,
			RegistrationDate: registrationDate,
			Address:          address,
			Phone:            phone,
		}
	}

	return c.JSON(http.StatusOK, response)
}
func handleUpdate(c echo.Context) error {

	var updateRequest UpdateRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	// Execute the update query against the database
	db, err := connectDB()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	// Execute the update query using parameters from the JSON payload
	_, err = db.Exec("UPDATE items SET "+updateRequest.Field+" = ? WHERE title = ? AND description = ?",
		updateRequest.Value, updateRequest.Title, updateRequest.Description)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to execute update query"})
	}

	// Return a success response
	return c.JSON(http.StatusOK, map[string]string{"message": "Update successful"})
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
	e.GET("/user", getUserDetails)
	e.POST("/adduser", adduseracc)
	e.POST("/additem", additems)
	e.POST("/updateitem", handleUpdate)

	e.Logger.Fatal(e.Start(":1323"))
}
