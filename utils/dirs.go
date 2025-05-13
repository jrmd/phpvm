/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package utils

import (
	"errors"
	"os"
	"path"
)

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func PhpVMPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".phpvm"), nil
}

func GetEnvDir() (string, error) {
	sess := os.Getenv("PHPVM_SESSION")

	if sess == "" {
		return PhpVMPath()
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".local/state/phpvm_multishell", sess), nil
}
