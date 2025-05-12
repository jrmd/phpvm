/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"github.com/jrmd/phpvm/utils"
	"github.com/spf13/cobra"
)

// cdCmd represents the test command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Use composer php requirements",
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetAppropriateVersion()
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
