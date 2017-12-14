package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "To get in touch",
	Long:  `Yoshi sent to you your information where you want`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SOURCE")
		fmt.Println(args)
		fmt.Println(cmd.Flags().Args())

		site, _ := cmd.Flags().GetBool("site")
		destination, _ := cmd.Flags().GetBool("destination")

		fmt.Println("site:", site)
		fmt.Println("destination:", destination)
	},
}

func init() {
	RootCmd.AddCommand(sourceCmd)

	// Exemple Commande : ./Yoshi source https://devTo.com --destination telegram --site google.com

	sourceCmd.Flags().BoolVarP("site", "s", false, "Done le site de news vous voulez")
	sourceCmd.Flags().BoolVarP("destination", "d", false, "Sur quelle plateforme souhaitez vous être informé")
}
