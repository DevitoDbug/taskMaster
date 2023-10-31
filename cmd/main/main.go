package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"taskMaster/pkg/models"
	"taskMaster/pkg/routes"
)

func main() {
	port := ":8000"

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("../../static/index.html"))
		tasks := []models.Task{
			{Id: 1, Name: "Task1", Completed: false},
			{Id: 2, Name: "Task2", Completed: true},
			{Id: 3, Name: "Task3", Completed: false},
		}
		tmp.Execute(w, tasks)

	}).Methods("GET")
	routes.RegisterRoutes(r)

	fmt.Println("Starting sever at port ", port)
	http.ListenAndServe(port, r)
}
