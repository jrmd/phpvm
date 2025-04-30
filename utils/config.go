/*
Copyright © 2025 Jerome Duncan <jerome@jrmd.dev>
*/
package utils

import (
	"encoding/json"
	"os"
	"path"
)

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
