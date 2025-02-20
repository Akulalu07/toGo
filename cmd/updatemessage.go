package cmd

import (
	"fmt"
	"strconv"
	"toGo/db"

	"github.com/spf13/cobra"
)

// updatemessageCmd represents the updatemessage command
var updatemessageCmd = &cobra.Command{
	Use:   "updatemessage",
	Short: "Update message in task or note",
	Long:  ``,
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
			fmt.Println("Please enter the ID")
			return
		}
		id, err := strconv.Atoi(args[0])
		message := args[1]
		if err != nil {
			fmt.Println("You need to provide a number, error: ", err)
			return
		}
		if notes {
			db.RemessNote(id, message)
		} else {
			db.RemessTask(id, message)
		}
	},
}

func init() {
	rootCmd.AddCommand(updatemessageCmd)
	updatemessageCmd.Flags().BoolP("tasks", "t", false, "Use this command for tasks")
	updatemessageCmd.Flags().BoolP("notes", "n", false, "Use this command for notes")
}
