package db

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Note represents the structure of a note with fields Id, Message, and Date.
type Note struct {
	Id      int    `gorm:"primaryKey"` // Unique identifier for the note
	Message string // Content of the note
	Date    string // Creation date of the note
}

// Task represents the structure of a task with fields Id, Message, Date, and Flag.
type Task struct {
	Id      int    `gorm:"primaryKey"` // Unique identifier for the task
	Message string // Content of the task
	Date    string // Creation date of the task
	Flag    bool   // Task completion status (completed or not)
}

// PrintNote prints all notes in a table format.
func PrintNote() {
	notes := GetNotes()                                // Retrieve all notes
	t := table.NewWriter()                             // Create a new table writer
	t.AppendHeader(table.Row{"Id", "Message", "Date"}) // Add table header

	// Add rows with notes to the table
	for _, note := range notes {
		t.AppendRow(table.Row{note.Id, note.Message, note.Date})
	}

	fmt.Println(t.Render()) // Print the table to the screen
}

// PrintTasks prints all tasks in a table format.
func PrintTasks() {
	tasks := GetTasks()                                        // Retrieve all tasks
	t := table.NewWriter()                                     // Create a new table writer
	t.AppendHeader(table.Row{"Id", "Message", "Date", "Flag"}) // Add table header

	// Add rows with tasks to the table
	for _, task := range tasks {
		fl := "❌" // By default, the task is not completed
		if task.Flag {
			fl = "✅" // If the task is completed, set the corresponding symbol
		}

		t.AppendRow(table.Row{task.Id, task.Message, task.Date, fl})
	}

	fmt.Println(t.Render()) // Print the table to the screen
}
