/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jrmd/phpvm/utils"
	"github.com/spf13/cobra"
)

type Env struct {
	Dir string
	UseOnCd bool
	MultiShell bool
	Now bool
	SessionId string
}

func GenZSH(env Env) {
		script := ""
	
		if env.UseOnCd {
			script += `
				function chpwd() {
						phpvm use &>/dev/null # run phpvm cd
				}
			`
		}

		if env.MultiShell {
			script += fmt.Sprintf(`
				export PHPVM_SESSION="%s"
			`, env.SessionId)
		}


		script += fmt.Sprintf(`
			export PATH="%s/bin:%s/sbin:$PATH"
		`, env.Dir, env.Dir)

		if env.Now {
			script += `
				phpvm use &>/dev/null
			`
		} else {
			script += `
				phpvm default &>/dev/null
			`
		}

		fmt.Println(script)

}


func GenBash(env Env) {
	script := ""
	
		if env.UseOnCd {
			script += `
				__phpvmcd() {{
					\cd "$@" || return $?
					phpvm use
				}}

				alias cd=__phpvmcd
			`
		}

		if env.MultiShell {
			script += fmt.Sprintf(`
				export PHPVM_SESSION="%s"
			`, env.SessionId)
		}


		script += fmt.Sprintf(`
			export PATH="%s/bin:%s/sbin:$PATH"
		`, env.Dir, env.Dir)

		if env.Now {
			script += `
				phpvm use &>/dev/null
			`
		} else {
			script += `
				phpvm default &>/dev/null
			`
		}

		fmt.Println(script)
}

func GenFish(env Env) {
		script := ""
	
		if env.UseOnCd {
			script += `
				function __phpvmoncd --on-event cd
					phpvm use > /dev/null ^ /dev/null
				end
			`
		}

		if env.MultiShell {
			script += fmt.Sprintf(`
				set -gx PHPVM_SESSION "%s"
			`, env.SessionId)
		}


		script += fmt.Sprintf(`
			set -gx PATH "%s/bin" "%s/sbin" $PATH
		`, env.Dir, env.Dir)

		if env.Now {
			script += `
				phpvm use > /dev/null ^ /dev/null
			`
		} else {
			script += `
				phpvm default > /dev/null ^ /dev/null
			`
		}

		fmt.Println(script)
}

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env [shell]",
	Short: "Print and set up required environment variables for",
	Run: func(cmd *cobra.Command, args []string) {
		useOnCd := cmd.Flag("use-on-cd").Value.String() == "true"
		multiShell := cmd.Flag("multi-shell").Value.String() == "true"
		now := cmd.Flag("multi-shell").Value.String() == "true"
		
		session := fmt.Sprintf("%d_%d", os.Getpid(), time.Now().UnixMicro())

		if multiShell {
			os.Setenv("PHPVM_SESSION", session)
		} else {
			os.Setenv("PHPVM_SESSION", "")
		}

		dir, err := utils.GetEnvDir()

		if err != nil {
			os.Exit(1)
		}
		env := Env {
			dir,
			useOnCd,
			multiShell,
			now,
			session,
		}

		switch args[0] {
			case "zsh":
				GenZSH(env)
			case "bash":
				GenBash(env)
			case "fish":
				GenFish(env)
		}
	},
}

func init() {
	rootCmd.AddCommand(envCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	envCmd.Flags().BoolP("use-on-cd", "c", false, "Use phpvm cd on cd")
	envCmd.Flags().BoolP("multi-shell", "m", false, "Different phpversion per shell")
	envCmd.Flags().BoolP("now", "n", false, "run phpvm use now")
}
