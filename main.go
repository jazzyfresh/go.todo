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
	fmt.Println("Go to http://localhost:8080 to get to it")
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/task", TaskPage)
	http.ListenAndServe(":8080", nil)
}

func RootPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "BOOYAH\n")
}

func TaskPage(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var task Task
	decoder.Decode(&task)
	fmt.Println(task)
	if (Task{}) != task {
		tasks = append(tasks, task)
	}
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
