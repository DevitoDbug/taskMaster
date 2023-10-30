package routes

import (
	"github.com/gorilla/mux"
	"taskMaster/pkg/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks/", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.GetTaskById).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.CreateTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("GET")
}
