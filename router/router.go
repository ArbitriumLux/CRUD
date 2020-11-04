package router

import (
	"Serverlox/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const port = ":9090"

func Router() {
	router := mux.NewRouter()
	controller := middleware.Controller{
		Field: 10,
	}
	//TODO add 'GET' 'PUT' 'POST' 'DELETE' methods
	router.HandleFunc("/", middleware.IndexPage).Methods("GET")
	router.HandleFunc("/customers", controller.GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", middleware.GetCustomer).Methods("GET")
	router.HandleFunc("/delete", middleware.DeleteCustomer)
	router.HandleFunc("/create/customer", middleware.CreateCustomer).Methods("POST")
	router.HandleFunc("/update", middleware.UpdateCustomer)
	router.HandleFunc("/edit", middleware.Edit)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))

}
