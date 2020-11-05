package middleware

import (
	"Serverlox/models"
	"Serverlox/server"
	_ "database/sql"
	"encoding/json"
	_ "encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"html/template"
	"log"
	_ "log"
	"net/http"  // used to access the request and response object of the api
	_ "strconv" // package used to covert string into int type

	"github.com/gorilla/mux"

	_ "github.com/joho/godotenv" //TODO
	_ "github.com/lib/pq"
)

type Controller struct {
	Field int
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

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	server.DataBaseConnection()
	//Loading HTML template and Error handling
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var customers []models.Customer
	params := mux.Vars(r)
	log.Print(customers)
	server.Db.First(&customers, params["id"])
	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer server.Db.Close()

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/customers", 301)
	log.Println("delete")
	server.DataBaseConnection()
	defer server.Db.Close()

	emp := r.URL.Query().Get("id")

	var customers []models.Customer
	server.Db.Delete(&customers, emp)
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	server.DataBaseConnection()
	var customers models.Customer
	if r.Method == "POST" {
		customers.FirstName = r.FormValue("FirstName")
		customers.LastName = r.FormValue("LastName")
		customers.Email = r.FormValue("Email")
		customers.Gender = r.FormValue("Gender")
		customers.Birthday = r.FormValue("Birthday")

	}
	defer server.Db.Close()
	json.NewDecoder(r.Body).Decode(&customers)
	server.Db.Create(&customers)
	http.Redirect(w, r, "/customers", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	server.DataBaseConnection()
	nId := r.URL.Query().Get("id")
	var customers models.Customer
	server.Db.First(&customers, nId)
	tmpl, err := template.ParseFiles("static/edit.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer server.Db.Close()
	http.Redirect(w, r, "/customers", 301)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	//................................
}
