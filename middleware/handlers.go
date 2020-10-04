package middleware

import (
	"Serverlox/models"
	"html/template"
	"net/http"
)

func UsersPage(w http.ResponseWriter, r *http.Request) {
	users := []models.User{{"ANTON", "Sargar ASS", false},
		{"Эдуард", "Педуардов", true}}
	tmpl, err := template.ParseFiles("static/users.html")

	//Error handling
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func IndexPage(w http.ResponseWriter, r * http.Request){
	tmpl, err := template.ParseFiles("static/index.html")

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
