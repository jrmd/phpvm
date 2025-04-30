/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package utils

import (
	"errors"
	"os"
	"os/exec"
	"path"
)

func VersionPath(version string) string {
	return "/opt/homebrew/opt/php@" + version
}

func VersionExists(version string) bool {
	exists, _ := FileExists(VersionPath(version))
	return exists
}

func SetVersion(version string) error {
	if !VersionExists(version) {
		return errors.New("version does not exist")
	}

	src := VersionPath(version)

	target, err := PhpVMPath()

	if err != nil {
		return err
	}
	os.MkdirAll(target, 0755)

	dirs := []string{"bin", "sbin"}

	for _, dir := range dirs {
		srcDir := path.Join(src, dir)
		targetDir := path.Join(target, dir)
		exec.Command("ln", "-sfn", srcDir, targetDir).Run()
	}

	return nil
}
