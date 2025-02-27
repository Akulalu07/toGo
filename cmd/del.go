package cmd

import (
	"strconv"
	"toGo/db"
	"toGo/utils"

	"github.com/spf13/cobra"
)

// delCmd represents the delete command
var delCmd = &cobra.Command{
	Use:   "del [id]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a note or task by ID",
	Long: `The del command allows you to delete notes or tasks from your list. 
You can specify what you want to delete using the --tasks or --notes flags. 
This command requires the ID of the item you want to delete.

Usage examples:
- Delete a task: toGo del --tasks 1
- Delete a note: toGo del --notes 2

Please note that you can only use one of the flags 
(tasks or notes) at a time. If you do not specify the ID, the command will return an error.`,
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
			utils.Fatal("Please enter the ID")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			utils.Fatal("You need to provide a number, error: ", err)
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
	delCmd.Flags().BoolP("tasks", "t", false, "Use this command for tasks")
	delCmd.Flags().BoolP("notes", "n", false, "Use this command for notes")
}
