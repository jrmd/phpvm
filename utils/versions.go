/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package utils

import (
	"encoding/json"
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

	target, err := GetEnvDir()

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

	WriteCurrent(version)

	return nil
}

type ShellConf struct {
	Current string `json:"current"`
}

func (s ShellConf) Write() {
	configDir, err := GetEnvDir()
	if err != nil {
		return
	}

	configFile := path.Join(configDir, "config.json")
	json, _ := json.Marshal(s)

	os.WriteFile(configFile, json, 0644)
}

func ShellConfigExists() (bool, error) {
	envDir, err := GetEnvDir()

	if err != nil {
		return false, err
	}

	configFile := path.Join(envDir, "config.json")

	if ok, err := FileExists(configFile); !ok || err != nil {
		return false, err
	}

	return true, nil
}


func WriteCurrent( version string ) {
	shell := GetShell()
	shell.Current = version
	shell.Write()
}

func GetCurrent() string {
	shell := GetShell()
	return shell.Current
}

func GetShell() ShellConf {
	shellConf := ShellConf{}

	if exists, _ := ShellConfigExists(); !exists {
		return shellConf
	}

	configDir, err := GetEnvDir()
	if err != nil {
		return shellConf
	}
	configFile := path.Join(configDir, "config.json")

	
	conf, err := os.ReadFile(configFile)

	if err != nil {
		return shellConf
	}

	_ = json.Unmarshal(conf, &shellConf)
	return shellConf

}
