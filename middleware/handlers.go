package middleware

import (
	"Serverlox/models"
	"Serverlox/server"
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"html/template"
	"log"
	_ "log"
	"net/http" // used to access the request and response object of the api
	"strconv"
	_ "strconv" // package used to covert string into int type

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Controller struct {
	Field int
}

func (c *Controller) GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c.Field)
	server.DataBaseConnection()
	//Loading HTML template and Error handling
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var customers []models.Customer
	server.Db.Find(&customers)
	//json.NewEncoder(w).Encode(&customers)
	//pass values from DataBase (customers var)
	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer server.Db.Close()
}

// CreateCustomer create a customer in the postgres db
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty customer of type models.Customer
	var customer models.Customer
	// decode the json request to customer
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call insert customer function and pass the customer
	insertID := insertCustomer(customer)
	// format a response object
	res := response{
		ID:      insertID,
		Message: "Customer created successfully",
	}
	//send the response
	json.NewEncoder(w).Encode(res)
}

// GetCustomer will return a single customer by its id
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the customerid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	// call the getCustomer function with customerid to retrieve a single customer
	customer, err := getCustomer(int64(id))

	if err != nil {
		log.Fatalf("Unable to get customer. %v", err)
	}
	// send the response
	json.NewEncoder(w).Encode(customer)
}

// GetAllCustomer will return all the customers
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the customers in the db
	customers, err := getAllCustomers()

	if err != nil {
		log.Fatalf("Unable to get all customer. %v", err)
	}

	// send all the customers as response
	json.NewEncoder(w).Encode(customers)
}

// UpdateCustomer update customer's detail in the postgres db
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the customerid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty customer of type models.Customer
	var customer models.Customer

	// decode the json request to customer
	err = json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update customer to update the customer
	updatedRows := updateCustomer(int64(id), customer)

	// format the message string
	msg := fmt.Sprintf("Customer updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteCustomer delete customer's detail in the postgres db
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the customerid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteCustomer, convert the int to int64
	deletedRows := deleteCustomer(int64(id))

	// format the message string
	msg := fmt.Sprintf("Customer updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")

	// Error handling
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

//------------------------- handler functions ----------------
// insert one customer in the DB
func insertCustomer(customer models.Customer) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO customers (FirstName, LastName, Birthday, Gender, Email, Adress) VALUES ($1, $2, $3, $4, $5, $6) RETURNING customerid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, customer.FirstName, customer.LastName, customer.Birthday, customer.Gender, customer.Email, customer.Adress).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one customer from the DB by its customerid
func getCustomer(id int64) (models.Customer, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a customer of models.Customer type
	var customer models.Customer

	// create the select sql query
	sqlStatement := `SELECT * FROM customers WHERE customerid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to customer
	err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Birthday, &customer.Gender, &customer.Email, &customer.Adress)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return customer, nil
	case nil:
		return customer, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty customer on error
	return customer, err
}

// get one user from the DB by its customerid
func getAllCustomers() ([]models.Customer, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var customers []models.Customer

	// create the select sql query
	sqlStatement := `SELECT * FROM customers`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var customer models.Customer

		// unmarshal the row object to customer
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Birthday, &customer.Gender, &customer.Email, &customer.Adress)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the customer in the customers slice
		customers = append(customers, customer)

	}

	// return empty customer on error
	return customers, err
}

// update customer in the DB
func updateCustomer(id int64, customer models.Customer) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE customers SET FirstName=$2, LastName=$3, Birthday=$4, Gender=$5, Email=$6, Adress=$7 WHERE customerid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, customer.FirstName, customer.LastName, customer.Birthday, customer.Gender, customer.Email, customer.Adress)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete customer in the DB
func deleteCustomer(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM customers WHERE customerid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
