package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Watch Together v0.2 Alpha")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
