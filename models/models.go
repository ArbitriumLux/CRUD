package models

import "github.com/jinzhu/gorm"

type User struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	IsFired bool
}

//This struct will be implemented in this project !!!!!
type Customer struct {
	gorm.Model
	FirstName string 	`json:"first_name"`
	LastName string		`json:"last_name"`
	Birthday string		`json:"birthday"`
	Gender string		`json:"gender"`
	Email string		`json:"email"`
	Address string		`json:"address"`
}