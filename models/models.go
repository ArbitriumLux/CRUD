package models

type User struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	IsFired bool
}