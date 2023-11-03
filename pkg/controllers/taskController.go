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

var tasks []models.Task
var html string = `
  <div id="divId%v" class="h-full flex flex-row justify-between border-b-2 border-gray-300 p-2">
    <span {{if %v}} class="text-gray-500" {{end}}>%v</span>
    <div class="h-full flex gap-1 self-end text-gray-50">
      <button
        {{if %v}}
        class="bg-purple-400 text-gray-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
        {{else}}
        class="bg-gray-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
        {{end}}
		hx-trigger="click"
        hx-get="/tasks/%v/complete"
        hx-target="#divId%v"
        hx-swap="outerHTML"
      >
        {{if %v}}Complete{{else}}Pending{{end}}
      </button>
      <button
	  		{{if %v}}
            class="bg-purple-500 w-16 px-1 h-6 text-xs rounded-lg self-end"
            {{else}}
            class="bg-purple-400 w-16 px-1 h-6 text-gray-500 text-xs rounded-lg self-end"
            {{end}}
			hx-trigger="click"
			hx-delete="/tasks/%v/Delete"
			hx-target="#divId%v"
			hx-swap="outerHTML"
      >
        Delete
      </button>
    </div>
  </div>`




func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("static/index.html"))

	tmp.Execute(w, tasks)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	var idIndex int
	ID, err := utility.GetTaskID(r)
	if err != nil {
		http.Error(w, "Failed to retrieve task ID", http.StatusBadRequest)
		return
	}

	//toggle completed
	for index, task := range tasks {
		if task.ID == ID {
			tasks[index].Completed = !task.Completed
			idIndex = index
			break
		}
	}
	temp := template.Must(template.New("div" + strconv.Itoa(ID)).Parse(fmt.Sprintf(html, ID, tasks[idIndex].Completed, tasks[idIndex].Name, tasks[idIndex].Completed, ID, ID, tasks[idIndex].Completed, tasks[idIndex].Completed, ID, ID)))

	err = temp.Execute(w, nil)
	if err != nil {
		log.Print(err)
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}
	taskID := len(tasks)
	taskName := r.FormValue("taskName")
	task := models.CreateTask(taskID, taskName)
	tasks = append(tasks, *task)

	temp := template.Must(template.New("div" + strconv.Itoa(taskID)).Parse(fmt.Sprintf(html, taskID, tasks[taskID].Completed, tasks[taskID].Name, tasks[taskID].Completed, taskID, taskID, tasks[taskID].Completed, tasks[taskID].Completed ,taskID, taskID)))
	temp.Execute(w, nil)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	ID, err := utility.GetTaskID(r)
	if err != nil {
		log.Print(ID)
		return
	}
	for index, task := range tasks {
		if task.ID == ID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
