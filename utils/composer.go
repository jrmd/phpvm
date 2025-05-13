package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
)

type ComposerConfig struct {
	Require struct {
		Php string `json:"php"`
	} `json:"require"`
}

func GetPHPFromComposer() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	composerPath := path.Join(wd, "composer.json")
	hasComposer, _ := FileExists(composerPath)

	if !hasComposer {
		return "", fmt.Errorf("no composer.json found in current directory")
	}

	data, err := os.ReadFile(composerPath)

	if err != nil {
		return "", err
	}

	composer := ComposerConfig{}

	err = json.Unmarshal(data, &composer)

	if err != nil {
		return "", err
	}

	if len(composer.Require.Php) == 0 {
		return "", fmt.Errorf("no PHP version found in composer.json")
	}

	return strings.ReplaceAll(composer.Require.Php, "^", ""), nil
}

func VersionMatches(version string) (bool, error) {
	composerConstraint, err := GetPHPFromComposer()
	if err != nil {
		return false, err
	}

	v, err := semver.NewVersion(version)
	if err != nil {
		log.Fatalf("Error parsing version: %s", err)
	}
	c, err := semver.NewConstraint(composerConstraint)
	if err != nil {
		log.Fatalf("Error parsing constraint: %s", err)
	}
	ok, errs := c.Validate(v)

	if len(errs) > 0 {
		return false, fmt.Errorf("error validating version: %v", errs)
	}

	return ok, nil
}

func AvailableVersions() semver.Collection {
	raw := GetConfig().Versions
	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			fmt.Printf("Error parsing available version: %s\b", err)
		}

		vs[i] = v
	}
	collection := semver.Collection(vs)
	sort.Sort(collection)
	slices.Reverse(collection)
	return collection

}

func GetAppropriateVersion() (string, error) {
	constraint, err := GetPHPFromComposer()

	if err != nil {
		return "", err
	}
	current := GetCurrent()

	if current != "" {
		if ok, _ := VersionMatches(current); ok {
			return "", nil
		}
	}

	versions := AvailableVersions()

	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return "", err
	}

	for _, v := range versions {
		if c.Check(v) {
			return v.Original(), nil
		}
	}

	return "", fmt.Errorf("no matching version found")
}

func SetAppropriateVersion() (bool, error) {
	version, err := GetAppropriateVersion()
	if err != nil {
		return false, err
	}

	if version == "" {
		return true, nil
	}

	err = SetVersion(version)

	if err == nil {
		config := GetConfig()
		config.SetCurrent(version)
		return true, nil
	}

	return false, nil
}
