package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"taskMaster/pkg/routes"
)

func main() {
	port := ":8000"

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	fmt.Println("Starting sever at port ", port)
	http.ListenAndServe(port, nil)
}
