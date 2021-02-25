// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/clivern/goenv/core/service"
)

// Golang type
type Golang struct {
	RootPath       string
	EnvironmentDir string
	ReleasesDir    string
	VersionFile    string
	FileSystem     *service.FileSystem
}

// NewGolangEnvironment creates a new instance
func NewGolangEnvironment(homePath string) *Golang {
	return &Golang{
		RootPath:       fmt.Sprintf("%s/%s", homePath, ".goenv"),
		ReleasesDir:    "releases",
		VersionFile:    ".go-version",
		EnvironmentDir: ".goenv",
		FileSystem:     &service.FileSystem{},
	}
}

// GetVersions returns a list of all available versions
func (g *Golang) GetVersions() []string {
	return golangReleases
}

// Install installs a golang version
func (g *Golang) Install(version string) error {
	return nil
}

// Uninstall removes a golang installed version
func (g *Golang) Uninstall(version string) error {

	path := fmt.Sprintf("%s/%s/%s", g.RootPath, g.ReleasesDir, version)

	if !g.FileSystem.DirExists(path) {
		return fmt.Errorf(
			"Unable to locate version %s path %s.",
			version,
			path,
		)
	}

	err := g.FileSystem.ClearDir(path)

	if err != nil {
		return fmt.Errorf(
			"Unable to clear version %s path %s.",
			version,
			path,
		)
	}

	err = g.FileSystem.DeleteDir(path)

	if err != nil {
		return fmt.Errorf(
			"Unable to delete version %s path %s.",
			version,
			path,
		)
	}

	return nil
}

// SetVersion sets the local or global golang version
func (g *Golang) SetVersion(path, version string) error {

	err := g.FileSystem.StoreFile(path, fmt.Sprintf("%s\n", version))

	if err != nil {
		return fmt.Errorf(
			"Unable to set go version to %s path %s.",
			version,
			path,
		)
	}

	return nil
}

// GetLocalVersion returns the local golang version
func (g *Golang) GetLocalVersion(dir string) (string, error) {
	var versionFile string

	baseDir := g.FileSystem.RemoveTrailingSlash(dir)

	for {
		if baseDir == "/" {
			versionFile = fmt.Sprintf("/%s", g.VersionFile)
		} else {
			versionFile = fmt.Sprintf("%s/%s", baseDir, g.VersionFile)
		}

		if g.FileSystem.FileExists(versionFile) {
			break
		}

		if baseDir == "/" {
			return "", fmt.Errorf("Unable to locate local version file")
		}

		baseDir = filepath.Dir(baseDir)
	}

	content, err := g.FileSystem.ReadFile(versionFile)

	if err != nil {
		return "", fmt.Errorf(
			"Unable to read local version file, path %s and error %s",
			versionFile,
			err.Error(),
		)
	}

	return strings.TrimSuffix(content, "\n"), nil
}

// GetGlobalVersion returns the global golang version
func (g *Golang) GetGlobalVersion() (string, error) {

	file := fmt.Sprintf("%s/%s", g.RootPath, g.VersionFile)

	if !g.FileSystem.FileExists(file) {
		return "", fmt.Errorf(
			"Global go version is not set yet, path %s.",
			file,
		)
	}

	content, err := g.FileSystem.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf(
			"Unable to read global version file, path %s and error %s",
			file,
			err.Error(),
		)
	}

	return strings.TrimSuffix(content, "\n"), nil
}

// Configure configures the environment
func (g *Golang) Configure() error {
	var err error

	if !g.FileSystem.DirExists(g.RootPath) {
		err = g.FileSystem.EnsureDir(g.RootPath, 755)
	}

	if !g.FileSystem.DirExists(fmt.Sprintf("%s/%s", g.RootPath, g.ReleasesDir)) {
		err = g.FileSystem.EnsureDir(fmt.Sprintf("%s/%s", g.RootPath, g.ReleasesDir), 755)
	}

	if err != nil {
		return fmt.Errorf("Unable to configure environment: %s", err.Error())
	}

	return nil
}

// CreateChim create a new shim
func (g *Golang) CreateShim() error {
	return nil
}

// Rehash gets a list of binaries under a certain
// go bin directory and create shim for them
func (g *Golang) Rehash() error {
	return nil
}
