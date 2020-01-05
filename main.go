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
	decoder := json.NewDecoder(request.Body)
	var task Task
	decoder.Decode(&task)

	// If task struct is not empty
	// https://stackoverflow.com/a/28447372/1766434
	if (Task{}) != task {
		if task.Uuid == "" {
			task.Uuid = uuid.New().String()
		}

		encoded, err := json.Marshal(task)
		if err != nil {
			log.Printf("Error marshalling task: %s\n", err)
		}

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(TASKS_BUCKET))
			err := b.Put([]byte(task.Uuid), []byte(encoded))
			fmt.Printf("Updated task: %s\n", task)
			return err
		})

	}
	tasksHtml := fmt.Sprintf("Created task: %s", task)
	fmt.Fprintf(writer, tasksHtml)
}

func RootPage(writer http.ResponseWriter, request *http.Request) {
	log.Println("Formatting tasks...")
	html := ""
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TASKS_BUCKET))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Uuid %s\n", k)
			var t Task
			err := json.Unmarshal(v, &t)
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

			html = html + fmt.Sprintf("- [%s] ", completedMarker)
			html = html + fmt.Sprintf("%s\n", t.Name)
		}
		fmt.Fprintf(writer, html)
		return nil
	})
}
