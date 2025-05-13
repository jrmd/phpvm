/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jrmd/phpvm/utils"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:     "use [version]",
	Short:   "Set a version of PHP to use",
	Long:    ``,
	Aliases: []string{"u"},
	Run: func(cmd *cobra.Command, args []string) {
		var version string
		var err error
		config := utils.GetConfig()
		if len(args) > 0 {
			version = args[0]
		} else {
			version, err = utils.GetAppropriateVersion()

			if err != nil {
				if config.Default == "" {
					os.Exit(1)
				}

				version = config.Default
			}
		}

		setDefault := cmd.Flag("default").Value.String() == "true"
		if !utils.VersionExists(version) {
			fmt.Printf("Version %s does not exist\n", version)
			os.Exit(1)
		}

		if err := utils.SetVersion(version); err != nil {
			fmt.Printf("Error setting version: %s\n%s", version, err)
			os.Exit(1)
		}

		fmt.Printf("Version %s set successfully\n", version)

		config.Versions = utils.UniqAppend(config.Versions, version)

		if setDefault {
			config.Default = version
		}

		config.Current = version
		config.Save()
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
