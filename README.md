# phpvm

`phpvm` is a simple command-line tool for managing multiple PHP versions on macOS, built with Go.

## Requirements

*   **macOS**: This tool is designed for macOS environments.
*   **Homebrew**: `phpvm` relies on Homebrew for installing PHP versions. Ensure you have Homebrew installed.

## Installation

Using go cli
```
go install github.com/jrmd/phpvm
```

Using brew
```
brew tap jrmd/tools
brew install phpvm
```

## Configuration

To make `phpvm` available in your terminal, you need to add its binary directory to your `PATH` and set up shell completion.

Add the following lines to your shell configuration file (e.g., `~/.zshrc`, `~/.bash_profile`, `~/.config/fish/config.fish`):

```bash
eval "$(phpvm completion <your_shell>)" # Replace <your_shell> with your shell (e.g., zsh, bash)
eval "$(phpvm env <your_shell>)"
```

After adding these lines, reload your shell configuration:

```bash
source ~/.zshrc # Or your shell config file
```


## Usage

```
phpvm [command]

Available Commands:
    cd          Use composer php requirements
    completion  Generate the autocompletion script for the specified shell
    default     Apply default version
    help        Help about any command
    install     Install a version of PHP
    use         Set a version of PHP to use
    Flags:
    -h, --help     help for phpvm
```


### `php env [shell]`

Sets up the env variables for the specified shell.

```
Print and set up required environment variables for bash, zsh, fish

Usage:
  phpvm env [shell] [flags]

Flags:
  -h, --help          help for env
  -m, --multi-shell   Different phpversion per shell
  -n, --now           run phpvm use now
  -c, --use-on-cd     Use phpvm cd on cd
```

e.g `phpvm env zsh --use-on-cd --multi-shell`
```
```

### `phpvm cd`

Takes the php value from `composer.json`'s `require` section and attempts to switch to a compatible PHP version.

```sh
$ phpvm cd --help
Takes the php value composer.json's require
Usage:
    phpvm cd [flags]
Flags:
    -h, --help   help for cd
```


### `phpvm completion`

Generate the autocompletion script for `phpvm` for the specified shell. See each sub-command's help for details on how to use the generated script.

```sh
$ phpvm completion --help

Generate the autocompletion script for phpvm for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
    phpvm completion [command]
    Available Commands:
    bash        Generate the autocompletion script for bash
    fish        Generate the autocompletion script for fish
    powershell  Generate the autocompletion script for powershell
    zsh         Generate the autocompletion script for zsh
Flags:
    -h, --help   help for completion
```


### `phpvm default`

Applies the default PHP version that was set with the `--default` flag during `install` or `use`.

```sh
$ phpvm default
Apply default version
Usage:
    phpvm default [flags]
Aliases:
    default, d
Flags:
    -h, --help   help for default
```


### `phpvm use`

Sets a specific version of PHP to be used in the current terminal session.

```sh
$ phpvm use --help

Set a version of PHP to use
Usage:
    phpvm use [version] [flags]
Aliases:
    use, u
Flags:
    -d, --default   Set as the version as the default version
    -h, --help      help for use
```


### `phpvm install`

Installs a specific version of PHP using Homebrew.

```sh
$ phpvm install --help
Install a version of PHP

Usage:
    phpvm install [version] [flags]
Aliases:
    install, i
Flags:
    -d, --default   Set as the version as the default version
    -h, --help      help for install
    -u, --use       Set as the version once installed
```

Example:

install 8.4 and set it as the default and set it to using.

```sh
phpvm install -ud 8.4
```
