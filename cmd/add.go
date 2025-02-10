package cmd

import (
	"fmt"
	"strings"
	"toGo/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [message]",
	Short: "Add a note or task",
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
			fmt.Println("Error retrieving tasks flag:", err)
			return
		}

		notes, err := cmd.Flags().GetBool("notes")
		if err != nil {
			fmt.Println("Error retrieving notes flag:", err)
			return
		}
		if notes == tasks {
			fmt.Println("Please use only one flag")
			return
		}
		if len(args) == 0 {
			fmt.Println("Please provide a message to create a note or task")
			return
		}
		message := strings.Join(args, " ")
		if notes {
			db.AddNotes(message)
		} else {
			db.AddTask(message)
		}
		fmt.Printf("Successfully added: %s\n", message)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("tasks", "t", false, "Use this command for tasks")
	addCmd.Flags().BoolP("notes", "n", false, "Use this command for notes")
}
