package main

import "time"

type Task struct {
	Uuid      string
	Name      string
	Completed bool

	//CreatedAt   time.Time
	//CompletedAt time.Time

	//Description string
	//Priority    int
	//Estimate    time.Duration

	//Recurrence  Recurrence

	// Dependants   []*Task // The parent depends on the child
	// Dependencies []*Task

	//Tags        []string
}

// calculate next action
// - Navigate graph
// - Priority weighting
// - Detect catch-22s (cycles)
// func (t *Task) getNextAction() *Task {
// 	for _, d := range t.Dependencies {
// 		if !d.Completed {
// 			return d
// 		}
// 	}
// 	return nil
// }

type Recurrence struct {
	frequency time.Duration
	endDate   time.Time
	quantity  int
}
