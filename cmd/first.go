package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

func init() {
	botCmd.AddCommand(first)
}

var first = &cobra.Command{
	Use:   "first",
	Short: "Search for projects repositories. Make required changes based on configuration in section in [apply-merge] section.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command to apply and merge")
		// gitlab.ReadRepoContents(viper.GetString("TOKEN"))
		fmt.Println("INSIDE FIRST COMMAND:", viper.GetString("TOKEN"))
	},
}
