package middleware

import (
	"Serverlox/models"
	_ "database/sql"
	_ "encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"html/template"
	_ "log"
	"net/http"  // used to access the request and response object of the api
	_ "strconv" // package used to covert string into int type
)
var db *gorm.DB

var err error

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	//Database connection
	db, err = gorm.Open("postgres","host=localhost port=5432 user=swkkd dbname=my_db sslmode=disable password=root")
	if err != nil{
		panic("failed to connect to database")
	}
	defer db.Close()

	//TODO this shouldn't be here!!!!!!!!!
	db.AutoMigrate(&models.Customer{})
	fmt.Println("migrated")

	//Loading HTML template and Error handling
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var customers []models.Customer
	db.Find(&customers)
	//json.NewEncoder(w).Encode(&customers)
	//pass values from DataBase (customers var)
	if err := tmpl.Execute(w, customers); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}



func IndexPage(w http.ResponseWriter, r * http.Request){
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
