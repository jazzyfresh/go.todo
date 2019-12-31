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

	// If task struct is not empty
	// https://stackoverflow.com/a/28447372/1766434
	if (Task{}) != task {
		// If task is already in list, Update it
		updated := false
		for i, t := range tasks {
			if t.Name == task.Name {
				tasks[i] = task
				updated = true
			}
		}
		// Otherwise, Add it
		if !updated {
			tasks = append(tasks, task)
		}
	}
	tasksHtml := makeTasksHtml()
	fmt.Fprintf(writer, tasksHtml)
}

func makeTasksHtml() string {
	html := ""
	for _, t := range tasks {

		// Because Go language doesn't have ternary operators
		// https://golang.org/doc/faq#Does_Go_have_a_ternary_form
		completedMarker := " "
		if t.Completed {
			completedMarker = "x"
		}

		html = html + fmt.Sprintf("- [%s] ", completedMarker)
		html = html + fmt.Sprintf("%s\n", t.Name)
	}
	return html
}
