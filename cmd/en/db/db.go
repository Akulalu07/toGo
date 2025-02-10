package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DbNote *gorm.DB // Database for notes
	DbTask *gorm.DB // Database for tasks
	err    error
	format = time.UnixDate // Date format for storage
)

// Dinit initializes the databases for notes and tasks, and performs schema migration.
func Dinit() {
	// Initialize the database for notes
	DbNote, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the notes database")
	}
	err = DbNote.AutoMigrate(&Note{}) // Automatic schema migration for notes
	if err != nil {
		panic("Failed to perform notes migration")
	}

	// Initialize the database for tasks
	DbTask, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the tasks database")
	}
	err = DbTask.AutoMigrate(&Task{}) // Automatic schema migration for tasks
	if err != nil {
		panic("Failed to perform tasks migration")
	}
}

// GetNotes returns all notes from the database.
func GetNotes() []Note {
	var notes []Note
	result := DbNote.Find(&notes)
	if result.Error != nil {
		fmt.Println("Error while retrieving notes:", result.Error)
	}
	return notes
}

// GetTasks returns all tasks from the database.
func GetTasks() []Task {
	var tasks []Task
	result := DbTask.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Error while retrieving tasks:", result.Error)
	}
	return tasks
}

// AddNotes adds a new note to the database.
func AddNotes(message string) {
	some := Note{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Set current date
	}
	result := DbNote.Create(&some)
	if result.Error != nil {
		fmt.Println("Error while adding note:", result.Error)
	}
}

// DelNotes deletes a note by the specified ID.
func DelNotes(id int) {
	result := DbNote.Delete(&Note{}, id)
	if result.Error != nil {
		fmt.Println("Error while deleting note:", result.Error)
	}
}

// AddTask adds a new task to the database.
func AddTask(message string) {
	some := Task{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Set current date
		Flag:    false,                           // Task not completed yet
	}
	result := DbTask.Create(&some)
	if result.Error != nil {
		fmt.Println("Error while adding task:", result.Error)
	}
}

// DelTask deletes a task by the specified ID.
func DelTask(id int) {
	result := DbTask.Delete(&Task{}, id)
	if result.Error != nil {
		fmt.Println("Error while deleting task:", result.Error)
	}
}

// DoneTask changes the status of a task to completed by the specified ID.
func DoneTask(id int) {
	var task Task
	result := DbTask.First(&task, id)
	if result.Error != nil {
		fmt.Println("Error while searching for task:", result.Error)
		return
	}
	task.Flag = true // Set the completed task flag
	result = DbTask.Save(&task)
	if result.Error != nil {
		fmt.Println("Error while updating task:", result.Error)
	}
}
