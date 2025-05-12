/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"log"

	"github.com/jrmd/phpvm/utils"
	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:     "default",
	Aliases: []string{"d"},
	Short:   "Apply default version",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if ok, err := utils.ConfigExists(); !ok || err != nil {
			log.Fatal("Config file does not exist")
		}

		config := utils.GetConfig()
		if config.Default == "" {
			log.Fatal("no default version set")
		}

		err := utils.SetVersion(config.Default)
		if err != nil {
			log.Fatal(err)
		}
		config.SetCurrent(config.Default)
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
