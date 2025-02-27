package cmd

import (
	"toGo/db"
	"toGo/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list -[n/t]",
	Short: "Get a list of notes or tasks",
	Long: `The list command allows you to retrieve a nicely formatted list 
of your tasks or notes. You can choose what you want to view 
by using the appropriate flags.

Usage examples:
- View all tasks: toGo list --tasks
- View all notes: toGo list --notes

Please note that you can only use one of the flags 
(tasks or notes) at a time. If both flags are specified, the command will return an error.`,
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
		if tasks && notes {
			utils.Fatal("Please use only one flag")
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
	listCmd.Flags().BoolP("tasks", "t", false, "Use this command for tasks")
	listCmd.Flags().BoolP("notes", "n", false, "Use this command for notes")
}
