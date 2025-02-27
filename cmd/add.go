package cmd

import (
	"fmt"
	"strings"
	"toGo/db"
	"toGo/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [message]",
	Short: "Add a note or task",
	Args:  cobra.ExactArgs(1),
	Long: `The add command allows you to add new notes or tasks to your list. 
You can specify what you want to add by using the appropriate flags.

Usage examples:
- Add a task: toGo add --tasks Buy milk
- Add a note: toGo add --notes Write project ideas

Please note that you must use only one of the flags 
(tasks or notes) at a time. If both flags are specified, the command will return an error. 
You also need to provide a message to create a note or task.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := cmd.Flags().GetBool("tasks")
		if err != nil {
			utils.Fatal("Error retrieving tasks flag:", err)
			return
		}

		notes, err := cmd.Flags().GetBool("notes")
		if err != nil {
			utils.Fatal("Error retrieving notes flag:", err)
			return
		}
		if notes == tasks {
			utils.Fatal("Please use only one flag")
			return
		}
		if len(args) == 0 {
			utils.Fatal("Please provide a message to create a note or task")
			return
		}
		message := strings.Join(args, " ")
		if notes {
			db.AddNotes(message)
		} else {
			db.AddTask(message)
		}
		utils.Good(fmt.Sprintf("Successfully added: %s", message))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("tasks", "t", false, "Use this command for tasks")
	addCmd.Flags().BoolP("notes", "n", false, "Use this command for notes")
}
