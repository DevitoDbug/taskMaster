package routes

import (
	"github.com/gorilla/mux"
	"taskMaster/pkg/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/tasks/", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}/complete", controllers.Complete).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.CreateTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("GET")
}
