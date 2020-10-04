package middleware

import (
	"Serverlox/models"
	"Serverlox/server"
	_ "database/sql"
	_ "encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"html/template"
	_ "log"
	"net/http"  // used to access the request and response object of the api
	_ "strconv" // package used to covert string into int type

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
