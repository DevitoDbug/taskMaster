package controllers

import (
	"html/template"
	"net/http"
	"taskMaster/pkg/models"
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

}

func GetTasks(w http.ResponseWriter, r *http.Request) {

}
func CreateTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
