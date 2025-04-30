/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"jrmd.dev/phpvm/utils"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a version of PHP",
	Aliases: []string{"i"},
	Long: ``,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
		use := cmd.Flag("use").Value.String() == "true"
		setDefault := cmd.Flag("default").Value.String() == "true"
		version := args[0]
		if utils.VersionExists(version) {
			if use {
				err := utils.SetVersion(version)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Printf("Version %s already installed\n", version)
			}
					if setDefault {
			config := utils.GetConfig()
			config.SetDefault(version)
		}

				os.Exit(0)
		}

		c := exec.Command("brew", "install", "php@"+version)
		if err := c.Run(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Version %s installed\n", version)
		if use {
			err := utils.SetVersion(version)
			if err != nil {
				log.Fatal(err)
			}
		}
		if setDefault {
			config := utils.GetConfig()
			config.SetDefault(version)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	installCmd.Flags().BoolP("use", "u", false, "Set as the version once installed")
	installCmd.Flags().BoolP("default", "d", false, "Set as the version as the default version")
}
