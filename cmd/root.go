package cmd

import "github.com/spf13/cobra"

// RootCmd is entry point of this cmd
var RootCmd = &cobra.Command{
	Use:   "yoshi",
	Short: "Yoshi is the niciest bot you ever seen.",
	Long:  `Yoshi is here to send to you some information, Tech news, weather.. and some personal stuff :)`,
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
