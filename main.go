package main

import (
	"Serverlox/models"
	"Serverlox/router"
	"Serverlox/server"
	"fmt"
)

// Basic execution logic: main > router.Router > handlers > models
func main() {
	//DataBase migrate!
	server.DataBaseConnection()
	server.Db.AutoMigrate(&models.Customer{})
	fmt.Println("Migration Successful!")
	defer server.Db.Close()

	//Start listening to http routes
	router.Router()

}
