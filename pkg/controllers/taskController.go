package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"taskMaster/pkg/models"
	"taskMaster/pkg/utility"
)

var tasks []models.Task

func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("../../static/index.html"))
	tasks = []models.Task{
		{ID: 1, Name: "Task1", Completed: false},
		{ID: 2, Name: "Task2", Completed: true},
		{ID: 3, Name: "Task3", Completed: false},
	}
	tmp.Execute(w, tasks)
}
func Complete(w http.ResponseWriter, r *http.Request) {
	ID, err := utility.GetTaskID(r)
	if err != nil {
		http.Error(w, "Failed to retrieve task ID", http.StatusBadRequest)
		return
	}

	// Check if the task ID is valid
	if ID < 0 || ID >= len(tasks) {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	//toggle completed
	tasks[ID].Completed = !tasks[ID].Completed

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task with ID " + strconv.Itoa(ID) + " has been completed"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

}
func CreateTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
