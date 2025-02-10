package cmd

import (
	"fmt"
	"strconv"
	"toGo/db"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Изменить статус задачи на выполнено",
	Long: `Команда done позволяет вам изменить статус задачи на "выполнено" по указанному ID. 
Это полезно для отслеживания прогресса выполнения задач в вашем списке.

Примеры использования:
- Отметить задачу как выполненную: toGo done 1

Обратите внимание, что необходимо указать ID задачи, которую вы хотите отметить как выполненную. 
Если ID не указан или указан неверно, команда выдаст ошибку.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Пожалуйста, введите ID задачи")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Необходимо указать число, ошибка: ", err)
			return
		}
		db.DoneTask(id)
		fmt.Printf("Задача с ID %d отмечена как выполненная.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
