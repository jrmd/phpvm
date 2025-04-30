/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"jrmd.dev/phpvm/utils"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
