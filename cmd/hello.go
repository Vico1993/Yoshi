package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Yoshi says to you hello",
	Long:  `This subcommand says hello`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello called")
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
}
