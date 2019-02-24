package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nocreep",
	Long:  `All software has versions. This is nocreep's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nocreep mobile analytics platform v0.9")
	},
}
