package cmd

import (
	"fmt"
	"strconv"
	"toGo/db"

	"github.com/spf13/cobra"
)

// delCmd представляет команду удаления
var delCmd = &cobra.Command{
	Use:   "del [id]",
	Short: "Удалить заметку или задачу по ID",
	Long: `Команда del позволяет вам удалять заметки или задачи из вашего списка. 
Вы можете указать, что хотите удалить, используя флаги --tasks или --notes. 
Эта команда требует указания ID элемента, который вы хотите удалить.

Примеры использования:
- Удалить задачу: toGo del --tasks 1
- Удалить заметку: toGo del --notes 2

Обратите внимание, что вы можете использовать только один из флагов 
(tasks или notes) за раз. Если вы не укажете ID, команда выдаст ошибку.`,
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
			fmt.Println("Пожалуйста, введите ID")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Необходимо указать число, ошибка: ", err)
			return
		}
		if notes {
			db.DelNotes(id)
		} else {
			db.DelTask(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().BoolP("tasks", "t", false, "Использовать эту команду для задач")
	delCmd.Flags().BoolP("notes", "n", false, "Использовать эту команду для заметок")
}
