package main

import (
	"Serverlox/router"
	"fmt"
)


func main() {
	// Basic execution logic: main > router.Router > handlers > models
	router.Router()
	//Logging errors

	fmt.Println("")
}
