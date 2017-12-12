package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "To get in touch",
	Long:  `Yoshi sent to you, depande on the source you want`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	RootCmd.AddCommand(sourceCmd)

	sourceCmd.PersistentFlags().String("site", "devTo", "Done le site de news vous voulez")
	// sourceCmd.Flags().BoolP("site", "s", false, "Done le site de news vous voulez")
}
