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

type Config struct {
	Default string `json:"default"`
}

func (c *Config) GetDefault() string {
	return c.Default
}

func (c *Config) SetDefault(version string) {
	c.Default = version
	WriteConfig(*c)
}

func ConfigExists() (bool, error) {
	configDir, err := PhpVMPath()
	if err != nil {
		return false, err
	}
	configFile := path.Join(configDir, "config.json")

	if ok, err := FileExists(configFile); !ok || err != nil {
		return false, err
	}

	return true, nil
}

func GetConfig() Config {
	cfg := Config{}

	configDir, err := PhpVMPath()
	if err != nil {
		return cfg
	}
	configFile := path.Join(configDir, "config.json")

	if ok, err := ConfigExists(); !ok || err != nil {
		return cfg
	}

	conf, err := os.ReadFile(configFile)

	if err != nil {
		return cfg
	}

	_ = json.Unmarshal(conf, &cfg)
	return cfg
}

func WriteConfig(cfg Config) error {
	configDir, err := PhpVMPath()
	if err != nil {
		return err
	}

	configFile := path.Join(configDir, "config.json")
	json, _ := json.Marshal(cfg)

	return os.WriteFile(configFile, json, 0644)
}

func PhpVMPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".phpvm"), nil
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
