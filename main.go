package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

// start the server
var tasks = []Task{}
var db *bolt.DB

const TASKS_BUCKET = "Tasks"

func main() {
	log.Println("Starting go.todo server")

	log.Println("Starting boltdb")
	dbTmp, err := bolt.Open("todo.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	db = dbTmp
	defer db.Close()

	// Create Tasks bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TASKS_BUCKET))
		if err != nil {
			return fmt.Errorf("Error creating bucket: %s", err)
		}
		return nil
	})

	log.Println("Go to http://localhost:8080 to get to it")
	http.HandleFunc("/", RootPage)
	http.HandleFunc("/task", TaskPage)
	http.ListenAndServe(":8080", nil)
}

func TaskPage(writer http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("Error updating task")
	decoder := json.NewDecoder(request.Body)
	var task Task
	decoder.Decode(&task)

	// If task struct is not empty https://stackoverflow.com/a/28447372/1766434
	if (Task{}) != task {
		if task.Uuid == "" {
			task.Uuid = uuid.New().String()
		}

		encoded, err := json.Marshal(task)
		if err != nil {
			message = fmt.Sprintf("Error marshalling task: %#+v\n", err)
		}

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(TASKS_BUCKET))
			err := b.Put([]byte(task.Uuid), []byte(encoded))
			message = fmt.Sprintf("Updated task: %#+v\n", task)
			return err
		})

	}
	log.Printf(message)
	fmt.Fprintf(writer, message)
}

func RootPage(writer http.ResponseWriter, request *http.Request) {
	log.Println("Formatting tasks...")
	html := ""

	// Begin boltdb transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASKS_BUCKET))
		c := b.Cursor()

		// Iterate over key/value pairs in boltdb
		for k, v := c.First(); k != nil; k, v = c.Next() {
			// Key is uuid
			// Log to server console
			// TODO: don't need to log every uuid for every read transaction
			fmt.Printf("Uuid %s\n", k)

			// Value is task data
			var t Task
			// Have to unmarshal the data into a json object
			err := json.Unmarshal(v, &t)
			// Hopefully that doesnt bork
			if err != nil {
				log.Println(err)
				return err
			}

			// Because Go language doesn't have ternary operators
			// https://golang.org/doc/faq#Does_Go_have_a_ternary_form
			completedMarker := " "
			if t.Completed {
				completedMarker = "x"
			}

			// Display task name & completed status, eg
			// - [x] task is complete
			html = html + fmt.Sprintf("- [%s] ", completedMarker)
			html = html + fmt.Sprintf("%s\n", t.Name)
		}
		fmt.Fprintf(writer, html)
		return nil
	})
}
