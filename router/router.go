package router

import (
	"Serverlox/middleware"
	"Serverlox/server"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)


func Router()  {
	router := mux.NewRouter()

	//TODO add 'GET' 'PUT' 'POST' 'DELETE' methods
	router.HandleFunc("/", middleware.IndexPage).Methods("GET")
	router.HandleFunc("/customers", middleware.GetCustomers).Methods("GET")
	//router.HandleFunc("/customer/{id}", middleware.GetCustomer).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(server.Port, handler))

}
