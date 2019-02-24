package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nocreep",
	Short: "nocreep is a lightweitgh mobile analytics with privacy in mind",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		s := ServerCommand{
			Store: StoreGroup{Type: "bolt"},
			Port:  3000}
		s.Execute()
	},
}

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
