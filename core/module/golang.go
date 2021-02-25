// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

// Golang type
type Golang struct {
	RootPath string
}

// NewGolangEnvironment creates a new instance
func NewGolangEnvironment() *Golang {
	return &Golang{
		RootPath: "/.goenv",
	}
}

// GetVersions returns a list of all available versions
func (g *Golang) GetVersions() []string {
	return golangReleases
}

// Install installs a golang version
func (g *Golang) Install(version string) (bool, error) {
	return true, nil
}

// Uninstall removes a golang installed version
func (g *Golang) Uninstall(version string) (bool, error) {
	return true, nil
}

// SetLocalVersion sets the local golang version
func (g *Golang) SetLocalVersion(version string) (bool, error) {
	return true, nil
}

// SetGlobalVersion sets the global golang version
func (g *Golang) SetGlobalVersion(version string) (bool, error) {
	return true, nil
}

// GetLocalVersion returns the local golang version
func (g *Golang) GetLocalVersion() (string, error) {
	return "", nil
}

// GetGlobalVersion returns the global golang version
func (g *Golang) GetGlobalVersion() (string, error) {
	return "", nil
}

// CreateChim create a new chim
func (g *Golang) CreateChim() (bool, error) {
	return true, nil
}

// Config configures the environment
func (g *Golang) Config() (bool, error) {
	return true, nil
}
