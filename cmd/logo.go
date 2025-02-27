package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logoCmd represents the logo command
var logoCmd = &cobra.Command{
	Use:   "logo",
	Short: "Displays the application logo",
	Long:  `This command displays the logo of the application in ASCII art.`,
	Run: func(cmd *cobra.Command, args []string) {
		printLogo()
	},
}

func printLogo() {
	fmt.Println(`
      _____                   _______                   _____                   _______         
     /\    \                 /::\    \                 /\    \                 /::\    \        
    /::\    \               /::::\    \               /::\    \               /::::\    \       
    \:::\    \             /::::::\    \             /::::\    \             /::::::\    \      
     \:::\    \           /::::::::\    \           /::::::\    \           /::::::::\    \     
      \:::\    \         /:::/~~\:::\    \         /:::/\:::\    \         /:::/~~\:::\    \    
       \:::\    \       /:::/    \:::\    \       /:::/  \:::\    \       /:::/    \:::\    \   
       /::::\    \     /:::/    / \:::\    \     /:::/    \:::\    \     /:::/    / \:::\    \  
      /::::::\    \   /:::/____/   \:::\____\   /:::/    / \:::\    \   /:::/____/   \:::\____\ 
     /:::/\:::\    \ |:::|    |     |:::|    | /:::/    /   \:::\ ___\ |:::|    |     |:::|    |
    /:::/  \:::\____\|:::|____|     |:::|    |/:::/____/  ___\:::|    ||:::|____|     |:::|    |
   /:::/    \::/    / \:::\    \   /:::/    / \:::\    \ /\  /:::|____| \:::\    \   /:::/    / 
  /:::/    / \/____/   \:::\    \ /:::/    /   \:::\    /::\ \::/    /   \:::\    \ /:::/    /  
 /:::/    /             \:::\    /:::/    /     \:::\   \:::\ \/____/     \:::\    /:::/    /   
/:::/    /               \:::\__/:::/    /       \:::\   \:::\____\        \:::\__/:::/    /    
\::/    /                 \::::::::/    /         \:::\  /:::/    /         \::::::::/    /     
 \/____/                   \::::::/    /           \:::\/:::/    /           \::::::/    /      
                            \::::/    /             \::::::/    /             \::::/    /       
                             \::/____/               \::::/    /               \::/____/        
                              ~~                      \::/____/                 ~~              
                                                       ~~                                               
	`)
}

func init() {
	rootCmd.AddCommand(logoCmd)
}
