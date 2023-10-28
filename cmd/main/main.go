package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"taskMaster/pkg/routes"
)

func main() {
	port := ":8000"

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("../../static/index.html"))
		tmp.Execute(w, nil)
	}).Methods("GET")
	routes.RegisterRoutes(r)

	fmt.Println("Starting sever at port ", port)
	http.ListenAndServe(port, r)
}
