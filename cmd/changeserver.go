package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"toGo/utils"

	"github.com/spf13/cobra"
)

// changeserverCmd represents the changeserver command
var changeserverCmd = &cobra.Command{
	Use:   "changeserver [server]",
	Short: "Change the server",
	Args:  cobra.ExactArgs(1),
	Long:  "Command to change the server that will be used in the application.",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetServer() // Load the current server configuration
		if err != nil {
			utils.Fatal(err, "Error loading server configuration:")
			return
		}

		if len(args) != 1 {
			utils.Fatal(nil, "Please input one argument") // Check if exactly one argument is provided
			return
		}
		news := strings.Join(args, " ") // Join the arguments into a single string

		// Check if the new server URL starts with http:// or https://
		if !strings.HasPrefix(news, "http://") && !strings.HasPrefix(news, "https://") {
			news = "http://" + news // Default to http:// if no prefix is provided
		}

		// Try to access the new server URL
		_, err = http.Get(news)
		if err != nil {
			// If the first attempt fails, try with https://
			news = "https://" + strings.TrimPrefix(news, "http://")
			_, err = http.Get(news)
			if err != nil {
				utils.Fatal(nil, "Please input a correct server") // If both attempts fail, print an error message
				return
			}
		}

		// Print the old server and the new server
		fmt.Printf("Changing server from: %s => %s\n", config, news)

		config = news           // Update the server configuration
		err = SetServer(config) // Save the updated configuration to the JSON file
		if err != nil {
			utils.Fatal(err, "Error saving server configuration:")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(changeserverCmd)
}
