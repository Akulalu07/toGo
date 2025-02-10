package cmd

import (
	"fmt"
	"strings"
	"toGo/db"

	"github.com/spf13/cobra"
)

// addCmd представляет команду добавления
var addCmd = &cobra.Command{
	Use:   "add [message]",
	Short: "Добавить заметку или задачу",
	Long: `Команда add позволяет вам добавлять новые заметки или задачи в ваш список. 
Вы можете указать, что именно хотите добавить, используя соответствующие флаги.

Примеры использования:
- Добавить задачу: toGo add --tasks Купить молоко
- Добавить заметку: toGo add --notes Записать идеи для проекта

Обратите внимание, что необходимо использовать только один из флагов 
(tasks или notes) за раз. Если оба флага указаны, команда выдаст ошибку. 
Также необходимо указать сообщение для создания заметки или задачи.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := cmd.Flags().GetBool("tasks")
		if err != nil {
			fmt.Println("Ошибка при получении флага tasks:", err)
			return
		}

		notes, err := cmd.Flags().GetBool("notes")
		if err != nil {
			fmt.Println("Ошибка при получении флага notes:", err)
			return
		}
		if notes == tasks {
			fmt.Println("Пожалуйста, используйте только один флаг")
			return
		}
		if len(args) == 0 {
			fmt.Println("Пожалуйста, укажите сообщение для создания заметки или задачи")
			return
		}
		message := strings.Join(args, " ")
		if notes {
			db.AddNotes(message)
		} else {
			db.AddTask(message)
		}
		fmt.Printf("Успешно добавлено: %s\n", message)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("tasks", "t", false, "Использовать эту команду для задач")
	addCmd.Flags().BoolP("notes", "n", false, "Использовать эту команду для заметок")
}
