package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"toGo/utils"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login [login] [password]",
	Short: "Login to the server",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		password := args[1]

		url, err := GetServer()
		if err != nil {
			utils.Fatal(err)
		}
		jsonData := []byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
		resp, err := http.Post(url+"/login", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			utils.Fatal("Error with request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			utils.Fatal("Error with read responce:", err)
			return
		}

		if string(body) == "{}" {
			SetCredentials(username, password)
			utils.Good(fmt.Sprintf("Log in server: %s with username: %s and password: %s\n", url, username, password))
		} else {
			utils.Fatal("Input correct username or password")
		}

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
