package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// start the server
var tasks = []Task{}

func main() {
	fmt.Println("Starting go.todo server")
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/task", TaskPage)
	http.ListenAndServe(":8080", nil)
}

func RootPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "BOOYAH")
}

func TaskPage(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var task Task
	decoder.Decode(&task)
	tasks = append(tasks, task)
	tasksHtml := makeTasksHtml()
	fmt.Fprintf(writer, tasksHtml)
}

func makeTasksHtml() string {
	html := ""
	for _, t := range tasks {
		html = html + fmt.Sprintf("* %s\n", t.Name)
	}
	return html
}

// accept task crud requests
