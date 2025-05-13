/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PHPVM version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
