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
	AddEnd_date     string `json:"end_date"`
	AddSeller       string `json:"Seller"`
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
type ItemDetails struct {
	Title             string `json:"Title"`
	Description       string `json:"Description"`
	Starting_bid      string `json:"Starting_bid"`
	Current_bid       string `json:"Current_bid"`
	End_Date          string `json:"End_Date"`
	CurrentBid_holder string `json:"CurrentBid_holder"`
	Seller            string `json:"Seller"`
	Category          string `json:"Category"`
}
type ItemsDetails struct {
	Items map[string]ItemDetails `json:"items"`
}
type UseraccountsDetails struct {
	Useraccounts map[string]UseraccountDetails `json:"useraccounts"`
}
type Useraccounts struct {
	Useraccounts map[string]Useraccount `json:"useraccounts"`
}
type Item struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Current_bid string `json:"Current_bid"`
	End_date    string `json:"end_date"`
	Seller      string `json:"Seller"`
}
type Items struct {
	Items map[string]Item `json:"items"`
}
type addbid struct {
	AddItem_number string `json:"item_number"`
	AddUser_id     string `json:"user_id"`
	AddBid_amount  string `json:"bid_amount"`
	AddBid_date    string `json:"bid_date"`
}
type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Field1      string `json:"field1"`
	Value1      string `json:"value1"`
	Field2      string `json:"field2"`
	Value2      string `json:"value2"`
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

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Shoes\"")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to execute query"})
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforclothing(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Clothing\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforelectronics(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Electronics\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforsports(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Sports\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforhealth(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Health\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforfurniture(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Furniture\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

	}
	return c.JSON(http.StatusOK, response)
}
func getiteminfoforother(c echo.Context) error {
	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	query, err := db.Query("SELECT Title, Description, Current_bid, End_date, Seller, Category FROM items WHERE category = \"Other\"")
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()

	response := Items{Items: map[string]Item{}}

	for query.Next() {

		var title string
		var description string
		var end_date string
		var current_bid string
		var seller string
		var category string

		err = (query).Scan(&title, &description, &current_bid, &end_date, &seller, &category)

		if err != nil {
			fmt.Println(err)
		}

		response.Items[title] = Item{Title: title, Description: description, Current_bid: current_bid, End_date: end_date, Seller: seller}

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
	result, err := db.Exec("INSERT INTO `nea_db`.`items` (`Title`, `Description`, `Starting_bid`, `Current_bid`, `End_date`, `Seller`, `Category`) VALUES (?, ?, ?, ?, ?, ?, ?);",
		requestData.AddTitle, requestData.AddDescription, requestData.AddStarting_bid, requestData.AddCurrent_bid, requestData.AddEnd_date, requestData.AddSeller, requestData.AddCategory)
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
func addbids(c echo.Context) error {
	// Parse request body
	var requestData addbid
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
	result, err := db.Exec("INSERT INTO `nea_db`.`bids` (`Item_number`, `User_id`, `Bid_amount`, `Bid_date`) VALUES (?, ?, ?, ?);",
		requestData.AddItem_number, requestData.AddUser_id, requestData.AddBid_amount, requestData.AddBid_date)
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

	requested_id := c.Param("email")
	fmt.Println(requested_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	query, err := db.Query("SELECT user_id, Firstname, Surname, email, Password, Registration_date, Address, phone FROM nea_db.useraccounts Where email = ? ", requested_id)
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
func getItemDetails(c echo.Context) error {

	db, err := sql.Open("mysql", "root:bball616.DAS@tcp(localhost:3306)/nea_db")

	requestedTitle := c.QueryParam("title")
	requestedDescription := c.QueryParam("description")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	querySQL := "SELECT Item_number, Title, Description, Starting_bid, Current_bid, End_date, CurrentBid_holder, Seller, Category FROM nea_db.items WHERE Title = ? AND Description = ?"
	query, err := db.Query(querySQL, requestedTitle, requestedDescription)
	fmt.Println("SQL Query:", querySQL, "Parameters:", requestedTitle, requestedDescription)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to execute query"})
	}
	defer query.Close()

	response := ItemsDetails{Items: map[string]ItemDetails{}}

	for query.Next() {
		var id string
		var title string
		var description string
		var starting_bid string
		var current_bid string
		var currentBid_holder string
		var end_date string
		var seller string
		var category string

		err = query.Scan(&id, &title, &description, &starting_bid, &current_bid, &end_date, &currentBid_holder, &seller, &category)
		if err != nil {
			fmt.Println(err)
		}

		response.Items[id] = ItemDetails{
			Title:             title,
			Description:       description,
			Starting_bid:      starting_bid,
			Current_bid:       current_bid,
			End_Date:          end_date,
			CurrentBid_holder: currentBid_holder,
			Seller:            seller,
			Category:          category,
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
	_, err = db.Exec("UPDATE items SET "+updateRequest.Field1+" = ?, "+updateRequest.Field2+" = ? WHERE title = ? AND description = ?", updateRequest.Value1, updateRequest.Value2, updateRequest.Title, updateRequest.Description)

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
	e.GET("/user/:email", getUserDetails)
	e.GET("/item", getItemDetails)
	e.POST("/adduser", adduseracc)
	e.POST("/additem", additems)
	e.POST("/updateitem", handleUpdate)
	e.POST("/addbid", addbids)

	e.Logger.Fatal(e.Start(":1323"))
}
