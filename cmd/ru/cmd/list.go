package cmd

import (
	"fmt"
	"toGo/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Получить список заметок или задач",
	Long: `Команда list позволяет вам получить красиво отформатированный список 
ваших задач или заметок. Вы можете выбрать, что именно хотите просмотреть, 
используя соответствующие флаги.

Примеры использования:
- Просмотреть все задачи: toGo list --tasks
- Просмотреть все заметки: toGo list --notes

Обратите внимание, что вы можете использовать только один из флагов 
(tasks или notes) за раз. Если оба флага указаны, команда выдаст ошибку.`,
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
		if tasks && notes {
			fmt.Println("Пожалуйста, используйте только один флаг")
			return
		}
		if tasks || (len(args) > 0 && args[0][0] == 't') {
			db.PrintTasks()
			return
		}
		if notes || (len(args) > 0 && args[0][0] == 'n') {
			db.PrintNote()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("tasks", "t", false, "Использовать эту команду для задач")
	listCmd.Flags().BoolP("notes", "n", false, "Использовать эту команду для заметок")
}
