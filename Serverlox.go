package main

import (
	"Serverlox/server"
	"Serverlox/handlers"
	"log"
	"net/http"
)


func main() {
	Router()
	println("Server listen on port"+ server.Port)
	err := http.ListenAndServe(server.Port, nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func Router()  {
	http.HandleFunc("/", handlers.IndexPage)
	http.HandleFunc("/users", handlers.UsersPage)
}