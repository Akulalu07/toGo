package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toGo",
	Short: "Manage your tasks with a simple CLI",
	Long: `toGo CLI is a convenient application for task management, 
allowing you to add, delete, and view your tasks 
directly from the command line. With toGo, you can easily organize 
your day, track progress, and stay productive.

Usage examples:
- Add a task: toGo add Buy milk
- Delete a task: toGo del 1
- View all tasks: toGo list

This application is built using the Cobra library for Go, 
which simplifies the creation of powerful CLI applications.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
