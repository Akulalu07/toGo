package db

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Note представляет структуру заметки с полями Id, Message и Date.
type Note struct {
	Id      int    `gorm:"primaryKey"` // Уникальный идентификатор заметки
	Message string // Содержимое заметки
	Date    string // Дата создания заметки
}

// Task представляет структуру задачи с полями Id, Message, Date и Flag.
type Task struct {
	Id      int    `gorm:"primaryKey"` // Уникальный идентификатор задачи
	Message string // Содержимое задачи
	Date    string // Дата создания задачи
	Flag    bool   // Статус выполнения задачи (выполнена или нет)
}

// PrintNote выводит все заметки в виде таблицы.
func PrintNote() {
	notes := GetNotes()                                // Получение всех заметок
	t := table.NewWriter()                             // Создание нового писателя таблицы
	t.AppendHeader(table.Row{"Id", "Message", "Date"}) // Добавление заголовка таблицы

	// Добавление строк с заметками в таблицу
	for _, note := range notes {
		t.AppendRow(table.Row{note.Id, note.Message, note.Date})
	}

	fmt.Println(t.Render()) // Вывод таблицы на экран
}

// PrintTasks выводит все задачи в виде таблицы.
func PrintTasks() {
	tasks := GetTasks()                                        // Получение всех задач
	t := table.NewWriter()                                     // Создание нового писателя таблицы
	t.AppendHeader(table.Row{"Id", "Message", "Date", "Flag"}) // Добавление заголовка таблицы

	// Добавление строк с задачами в таблицу
	for _, task := range tasks {
		fl := "❌" // По умолчанию задача не выполнена
		if task.Flag {
			fl = "✅" // Если задача выполнена, устанавливаем соответствующий символ
		}

		t.AppendRow(table.Row{task.Id, task.Message, task.Date, fl})
	}

	fmt.Println(t.Render()) // Вывод таблицы на экран
}
