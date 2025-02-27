package cmd

import (
	"fmt"
	"strconv"
	"toGo/db"
	"toGo/utils"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Args:  cobra.ExactArgs(1),
	Short: "Change the status of a task to completed",
	Long: `The done command allows you to change the status of a task to "completed" by the specified ID. 
This is useful for tracking the progress of tasks in your list.

Usage examples:
- Mark a task as completed: toGo done 1

Please note that you must specify the ID of the task you want to mark as completed. 
If the ID is not provided or is incorrect, the command will return an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utils.Fatal("Please enter the task ID")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			utils.Fatal("You need to provide a number, error: ", err)
			return
		}
		db.DoneTask(id)
		utils.Good(fmt.Sprintf("Task with ID %d has been marked as completed.", id))
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
