package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DbNote *gorm.DB // База данных для заметок
	DbTask *gorm.DB // База данных для задач
	err    error
	format = time.UnixDate // Формат даты для хранения
)

// Dinit инициализирует базы данных для заметок и задач, а также выполняет миграцию схемы.
func Dinit() {
	// Инициализация базы данных для заметок
	DbNote, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных для заметок")
	}
	err = DbNote.AutoMigrate(&Note{}) // Автоматическая миграция схемы для заметок
	if err != nil {
		panic("Не удалось выполнить миграцию заметок")
	}

	// Инициализация базы данных для задач
	DbTask, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных для задач")
	}
	err = DbTask.AutoMigrate(&Task{}) // Автоматическая миграция схемы для задач
	if err != nil {
		panic("Не удалось выполнить миграцию задач")
	}
}

// GetNotes возвращает все заметки из базы данных.
func GetNotes() []Note {
	var notes []Note
	result := DbNote.Find(&notes)
	if result.Error != nil {
		fmt.Println("Ошибка при получении заметок:", result.Error)
	}
	return notes
}

// GetTasks возвращает все задачи из базы данных.
func GetTasks() []Task {
	var tasks []Task
	result := DbTask.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Ошибка при получении задач:", result.Error)
	}
	return tasks
}

// AddNotes добавляет новую заметку в базу данных.
func AddNotes(message string) {
	some := Note{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Установка текущей даты
	}
	result := DbNote.Create(&some)
	if result.Error != nil {
		fmt.Println("Ошибка при добавлении заметки:", result.Error)
	}
}

// DelNotes удаляет заметку по указанному ID.
func DelNotes(id int) {
	result := DbNote.Delete(&Note{}, id)
	if result.Error != nil {
		fmt.Println("Ошибка при удалении заметки:", result.Error)
	}
}

// AddTask добавляет новую задачу в базу данных.
func AddTask(message string) {
	some := Task{
		Message: message,
		Date:    time.Now().UTC().Format(format), // Установка текущей даты
		Flag:    false,                           // Задача еще не выполнена
	}
	result := DbTask.Create(&some)
	if result.Error != nil {
		fmt.Println("Ошибка при добавлении задачи:", result.Error)
	}
}

// DelTask удаляет задачу по указанному ID.
func DelTask(id int) {
	result := DbTask.Delete(&Task{}, id)
	if result.Error != nil {
		fmt.Println("Ошибка при удалении задачи:", result.Error)
	}
}

// DoneTask изменяет статус задачи на выполнено по указанному ID.
func DoneTask(id int) {
	var task Task
	result := DbTask.First(&task, id)
	if result.Error != nil {
		fmt.Println("Ошибка при поиске задачи:", result.Error)
		return
	}
	task.Flag = true // Установка флага выполненной задачи
	result = DbTask.Save(&task)
	if result.Error != nil {
		fmt.Println("Ошибка при обновлении задачи:", result.Error)
	}
}
