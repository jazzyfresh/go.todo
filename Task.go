package main

import "time"

type Task struct {
	Name      string
	Completed bool
	//description string
	//parent      Task
	//children    []Task
	//tags        []string
	//recurrence  Recurrence
	//estimate    time.Duration
	//priority    int
	//createdAt   time.Time
	//completedAt time.Time
}

// calculate next action

type Recurrence struct {
	frequency time.Duration
	endDate   time.Time
	quantity  int
}
