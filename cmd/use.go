/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"jrmd.dev/phpvm/utils"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		setDefault := cmd.Flag("default").Value.String() == "true"
		if ! utils.VersionExists( version ) {
			fmt.Printf("Version %s does not exist\n", version)
			os.Exit(1)
		}

		if err := utils.SetVersion(version); err != nil {
			fmt.Printf("Error setting version: %s\n%s", version, err)
			os.Exit(1)
		}

		fmt.Printf("Version %s set successfully\n", version)
				if setDefault {
			config := utils.GetConfig()
			config.SetDefault(version)
		}

	},
}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolP("default", "d", false, "Set as the version as the default version")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
