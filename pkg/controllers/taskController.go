package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"taskMaster/pkg/models"
	"taskMaster/pkg/utility"
)

var tasks = []models.Task{
	{ID: 0, Name: "Task1", Completed: false},
	{ID: 1, Name: "Task2", Completed: true},
	{ID: 2, Name: "Task3", Completed: false},
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("../../static/index.html"))

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
	temp := template.Must(template.New("div" + strconv.Itoa(ID)).Parse(fmt.Sprintf(`
  <div id="divId%v" class="h-full flex flex-row justify-between border-b-2 border-gray-300 p-2">
    <span {{if %v}} class="text-gray-500" {{end}}>%v</span>
    <div class="h-full flex gap-1 self-end text-gray-50">
      <button
        {{if %v}}
        class="bg-purple-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
        {{else}}
        class="bg-gray-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
        {{end}}
		hx-trigger="click"
        hx-get="/tasks/%v/complete"
        hx-target="#divId%v"
        hx-swap="outerHTML"
      >
        {{if %v}}
        Complete
        {{else}}
        Pending
        {{end}}
      </button>
      <button
        class="bg-purple-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
      >
        Delete
      </button>
    </div>
  </div>`, ID, tasks[ID].Completed, tasks[ID].Name, tasks[ID].Completed, ID, ID, tasks[ID].Completed)))

	err = temp.Execute(w, nil)
	if err != nil {
		log.Print(err)
		return
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

}
func CreateTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
